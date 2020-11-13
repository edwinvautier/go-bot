package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/meteo"

	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
	// Shortening the import reference name seems to make it a bit easier
)

type GetWeather struct {
	analysis *wit.Analysis
	session  *discordgo.Session
	message  *discordgo.MessageCreate
}

func (command GetWeather) Execute() error {
	log.Info("GetWeather :", command.analysis)
	// Est ce que j'ai toutes les infos dont j'ai besoin dans analysis ?
	// non -> j'envoie un message : ou ça ?
	location := command.analysis.Location
	if len(location) == 0 {
		// message à l'utilisateur
		if command.message.Author.ID == command.session.State.User.ID {
			return nil
		}

		command.session.ChannelMessageSend(command.message.ChannelID, "Je sais pas compris, merci de reformuler ta votre demande")
		return errors.New("Could not retrieve location")
	}

	// oui je continue et je créé ma WeatherParams
	log.Info("You want the meteo in : ", location)

	weatherParams := meteo.WeatherParams{Location: location[0].Value}

	// Je renvoie un message
	meteo.FindWheatherByCity(&weatherParams)
	log.Info("channel", &command.session)
	if command.message.Author.ID == command.session.State.User.ID {
		return nil
	}
	command.session.ChannelMessageSend(command.message.ChannelID, "Voici la météo sur Moncul")

	// Potentiellement je renvoie des erreurs à plusieurs endroits
	return nil
}
