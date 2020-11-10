package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
	"errors"
)

// Dispatch a command depending on the analysis result we give
func Dispatch(a *wit.Analysis, s *discordgo.Session, m *discordgo.MessageCreate) error {
	if len(a.Intent) == 0 || a.Intent[0].Value == "" {
		return errors.New("Missing fields in analysis")
	}

	// Read the intent from the wit analysis result
	intentString := a.Intent[0].Value
	
	switch intentString {
	case "listen":
		QueryYoutubeVideo{analysis: a}.ExecuteYoutubeSearch(s, m)
	case "meteo":
		log.Info("You want the meteo")
	default: 
		return errors.New("Unknown command " + intentString)
	}

	return nil
}
