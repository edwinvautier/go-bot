package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/connectors"
)

// GenericCommand is the structure needed for every command.
// All Commands implements it.
type GenericCommand struct {
	Analysis *wit.Analysis
	Session  connectors.Discord
	Message  *discordgo.MessageCreate
}

// Build a command depending on the analysis result we give
func (gc *GenericCommand) Build() (Command, error) {
	if len(gc.Analysis.Intent) == 0 || gc.Analysis.Intent[0].Value == "" {
		return nil, errors.New("Missing fields in analysis")
	}

	// Read the intent from the wit analysis result
	intentString := gc.Analysis.Intent[0].Value

	switch intentString {
	case "listen":
		return QueryYoutubeVideoCommand{gc: gc}, nil
	case "meteo":
		return QueryWeatherCommand{gc: gc}, nil
	default:
		return QueryGoogleCommand{Connector: gc.Session, Message: gc.Message}, nil
	}

	return nil, nil
}
