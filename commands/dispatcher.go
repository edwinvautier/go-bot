package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

func Dispatch(a *wit.Analysis) {
	// Read the intent from the wit analysis result
	intentString := a.Intent[0].Value
	
	switch intentString {
	case "listen":
		music := a.Music[0].Value
		log.Info("You asked for music : ", music)
	case "meteo":
		location := a.Location[0].Value
		log.Info("You want the meteo in : ", location)
	}
}