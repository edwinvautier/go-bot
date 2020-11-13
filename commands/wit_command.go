package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
	log "github.com/sirupsen/logrus"
)

// AnalyzeSentence is the struct for our wit ai command
type AnalyzeSentence struct {
	Sentence string
}

// ExecuteWitCommand is the function that calls wit ai API to understand sentence
func (command AnalyzeSentence) ExecuteWitCommand() interface{} {
	analysis, err := wit.AnalyzeSentence(command.Sentence)
	if err != nil {
		log.Error("Could not retrieve analysis")
		return nil
	}

	return analysis
}
