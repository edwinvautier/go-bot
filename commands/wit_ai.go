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
	
	// Read informations from the analyzis interface returned by the wit command
	intentString := analysis.Intent[0].Value
	var value string

	if len(analysis.Location) != 0 {
		value = analysis.Location[0].Value
	} else if len(analysis.Music) != 0 {
		value = analysis.Music[0].Value
	}
	
	// Send a message to the user with the informations wit understood 
	command.Session.ChannelMessageSend(command.Message.ChannelID, "You want : " + intentString + "\n value : " + value)
}
