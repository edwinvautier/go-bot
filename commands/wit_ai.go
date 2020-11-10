package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
)

type AnalyzeSentence struct {
	Sentence 	string
	Session		*discordgo.Session
	Message		*discordgo.MessageCreate
}

func (command AnalyzeSentence) Execute() {
	analysis := wit.AnalyzeSentence(command.Sentence)
	
	intentString := analysis.Intent[0].Value
	var value string

	if len(analysis.Location) != 0 {
		value = analysis.Location[0].Value
	} else if len(analysis.Music) != 0 {
		value = analysis.Music[0].Value
	}
	
	command.Session.ChannelMessageSend(command.Message.ChannelID, "You want : " + intentString + "\n value : " + value)
}
	