package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/commands"
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

	analyzeCommand := commands.AnalyzeSentence{
		Sentence: sentence,
	}
	analysis := analyzeCommand.ExecuteWitCommand().(*wit.Analysis)
	if analysis == nil {
		s.ChannelMessageSend(m.ChannelID, "Pardon, je n'ai pas compris.")
		return
	}
	cmd, err := commands.Build(analysis, s, m)
	
	if err != nil {
		log.Error(err)
		s.ChannelMessageSend(m.ChannelID, "Echec lors de l'exécution de votre commande.")
		return
	}
	if cmd == nil {
		return
	}
	
	if err = cmd.Execute(); err != nil {
		s.ChannelMessageSend(m.ChannelID, "Echec lors de l'exécution de votre commande.")
		log.Error(err)
		return
	}
}
