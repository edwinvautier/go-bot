package main

import (
	"github.com/caarlos0/env"
	"github.com/edwinvautier/go-bot/conf"
	"github.com/edwinvautier/go-bot/discord"
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
	dg, err := discord.InitializeBot()
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
