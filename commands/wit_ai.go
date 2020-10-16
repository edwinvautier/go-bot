package commands

import (
	"github.com/edwinvautier/go-bot/apis/wit"
)

type AnalyzeSentence struct {
	Sentence 	string
}

func (command AnalyzeSentence) Execute() {
	wit.AnalyzeSentence(command.Sentence)
}
