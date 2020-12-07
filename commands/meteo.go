package commands

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/meteo"
	"github.com/edwinvautier/go-bot/connectors"

	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

type GetWeather struct {
	analysis  *wit.Analysis
	connector connectors.Discord
	message   *discordgo.MessageCreate
}

// Execute command to get results from api weather
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

	weatherParams := meteo.WeatherParams{Location: location[0].Value}


	wd := meteo.FindWheatherByCity(&weatherParams)
	log.Info(wd)
	_, err := command.connector.ChannelMessageSend(command.message.ChannelID, fmt.Sprintf("Voici la météo : %s", wd))
	if err != nil {
		log.Error("sendMessageErr: ", err)
	}
	return nil
}
