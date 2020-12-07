package commands

import (
	"errors"
	"fmt"

	"github.com/edwinvautier/go-bot/apis/meteo"
	log "github.com/sirupsen/logrus"
)

// QueryWeatherCommand is the command to ask for meteo
type QueryWeatherCommand struct {
	gc *GenericCommand
}

// Execute command to get results from api weather
func (command QueryWeatherCommand) Execute() error {
	location := command.gc.Analysis.Location

	if len(location) == 0 {
		wd, err := meteo.GetHereHandler()

		if err != nil {
			_, err = command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, "Je sais pas compris, vous n'exister pas")
			if err != nil {
				log.Error("sendMessageErr: ", err)
			}

			return errors.New("Could not retrieve location")
		}

		command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, fmt.Sprintf("Voici la météo : %s", wd))
		return nil
	}

	weatherParams := meteo.WeatherParams{Location: location[0].Value}

	wd := meteo.FindWheatherByCity(&weatherParams)
	_, err := command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, fmt.Sprintf("Voici la météo : %s", wd))
	if err != nil {
		log.Error("sendMessageErr: ", err)
	}
	return nil
}
