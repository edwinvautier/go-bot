package handlers

import (
	"github.com/edwinvautier/go-bot/commands"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"strings"
)

// MessageCreate is called when a new message is received by the bot
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, "assistant,") {
		return
	}
	sentence := strings.Split(m.Content, "assistant,")[1]
	
	analyzeCommand := commands.AnalyzeSentence {
		Sentence: sentence,
	}
	analysis := analyzeCommand.Execute().(*wit.Analysis)
	
	if err := commands.Dispatch(analysis); err != nil {
		log.Error(err)
	}
}
