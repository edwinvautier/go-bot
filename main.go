package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env"
	"github.com/edwinvautier/go-bot/conf"
	"github.com/edwinvautier/go-bot/handlers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database Setup
	dbcfg := conf.DbConfig{}
	if err := env.Parse(&dbcfg); err != nil {
		log.Fatal(err)
	}
	conf.InitializeDb(dbcfg.DbHost, dbcfg.DbUser, dbcfg.DbName, dbcfg.DbPort, dbcfg.DbPassword)
	conf.MakeMigrations()

	// Discord Bot initialization
	dg, err := initializeBot()
	if err != nil {
		log.Fatal("Error initializing discord API connection.")
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func initializeBot() (*discordgo.Session, error){
	discordToken, tokenExist := os.LookupEnv("DISCORD_TOKEN")
	if !tokenExist {
		log.Fatal("Missing environment variable : DISCORD_TOKEN")
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Error("Error creating Discord session, ", err)
		return nil, err
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handlers.MessageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error("Error opening discord connection, ", err)
		return nil, err
	}

	return dg, nil
}