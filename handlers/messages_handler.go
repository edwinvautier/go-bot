package handlers

import (
	"github.com/edwinvautier/go-bot/commands"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	analyzeCommand := commands.AnalyzeSentence {
		Sentence: m.Content,
	}
	analyzeCommand.Execute()
}
