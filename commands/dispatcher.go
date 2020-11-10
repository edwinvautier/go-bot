package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
	"errors"
)

// Dispatch a command depending on the analysis result we give
func Dispatch(a *wit.Analysis) error {
	if len(a.Intent) == 0 || a.Intent[0].Value == "" {
		return errors.New("Missing fields in analysis")
	}

	// Read the intent from the wit analysis result
	intentString := a.Intent[0].Value
	
	switch intentString {
	case "listen":
		log.Info("You asked for music")
	case "meteo":
		log.Info("You want the meteo")
	default: 
		return errors.New("Unknown command " + intentString)
	}

	return nil
}