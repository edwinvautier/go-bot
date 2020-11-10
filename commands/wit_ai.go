package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

// AnalyzeSentence is the struct for our wit ai command
type AnalyzeSentence struct {
	Sentence string
}

// Execute is the function implemented from the commandInterface
func (command AnalyzeSentence) Execute() interface{} {
	analysis := wit.AnalyzeSentence(command.Sentence)
	if nil == analysis {
		log.Error("Could not retrieve analysis")
	}

	return analysis
}
