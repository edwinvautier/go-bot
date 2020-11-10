package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

type AnalyzeSentence struct {
	Sentence 	string
}

func (command AnalyzeSentence) Execute() interface{} {
	analysis := wit.AnalyzeSentence(command.Sentence)
	if nil == analysis {
		log.Error("Could not retrieve analysis")
	}

	return analysis
}
