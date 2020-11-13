package commands

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/meteo"
	"github.com/edwinvautier/go-bot/connectors"

	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
	// Shortening the import reference name seems to make it a bit easier
)

type GetWeather struct {
	analysis  *wit.Analysis
	connector connectors.Discord
	message   *discordgo.MessageCreate
}

func (command GetWeather) Execute() error {
	location := command.analysis.Location

	if len(location) == 0 {
		wd, err := meteo.GetHereHandler()

		if err != nil {
			command.connector.ChannelMessageSend(command.message.ChannelID, "Je sais pas compris, vous n'exister pas")
			return errors.New("Could not retrieve location")
		}

		command.connector.ChannelMessageSend(command.message.ChannelID, fmt.Sprintf("Voici la météo : %s", wd))
		return nil
	}

	// oui je continue et je créé ma WeatherParams
	log.Info("You want the meteo in : ", location)

	weatherParams := meteo.WeatherParams{Location: location[0].Value}

	// Je renvoie un message
	wd := meteo.FindWheatherByCity(&weatherParams)
	log.Info(wd)
	command.connector.ChannelMessageSend(command.message.ChannelID, fmt.Sprintf("Voici la météo : %s", wd))

	// Potentiellement je renvoie des erreurs à plusieurs endroits
	return nil
}
