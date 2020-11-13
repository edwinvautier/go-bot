package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

// Build a command depending on the analysis result we give
func Build(a *wit.Analysis, s *discordgo.Session, m *discordgo.MessageCreate) (Command, error) {
	if len(a.Intent) == 0 || a.Intent[0].Value == "" {
		return nil, errors.New("Missing fields in analysis")
	}

	// Read the intent from the wit analysis result
	intentString := a.Intent[0].Value

	switch intentString {
	case "listen":
		return QueryYoutubeVideoCommand{analysis: a, session: s, message: m}, nil
	case "meteo":
		log.Info("You want the meteo")
	default:
		log.Error("unknown command: ", intentString)
		return nil, errors.New("Unknown command " + intentString)
	}

	return nil, nil
}
