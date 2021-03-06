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

	if !strings.HasPrefix(strings.ToLower(m.Content), "assistant,") {
		return
	}
	sentence := strings.Split(strings.ToLower(m.Content), "assistant,")[1]

	analyzeCommand := commands.AnalyzeSentence{
		Sentence: sentence,
	}
	var analysis *wit.Analysis

	// Call to wit to analyze sentence intents
	analysis = analyzeCommand.ExecuteWitCommand()
	if analysis == nil || len(analysis.Intent) < 1  {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pardon, je n'ai pas compris.")
		
		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		return
	}

	// If not enough confidence in user intent, ask google command
	if analysis.Intent[0].Confidence < 0.9 {
		askGoogle(s, m)
		return
	}

	// Else call command builder for the intent detected
	gc := commands.GenericCommand{Analysis: analysis, Session: s, Message: m}
	cmd, err := gc.Build()
	
	if err != nil {
		log.Error(err)
		_, err := s.ChannelMessageSend(m.ChannelID, "Je n'ai pas réussi à trouver ce qu'il vous fallait.")

		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		return
	}
	if cmd == nil {
		return
	}
	
	// If error during command creation, fallback solution is google command
	if err = cmd.Execute(); err != nil {
		askGoogle(s, m)
		return
	}
}

func askGoogle(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "Je n'ai pas très bien compris, je demande a google.")
	if err != nil {
		log.Error("sendMessageErr: ", err)
	}

	// Execute command to get results from google search browser
	googleCmd := commands.QueryGoogleCommand{Connector: s, Message: m}
	err = googleCmd.Execute()

	// Handle eventual errors
	if err != nil {
		log.Error(err)
		_, err := s.ChannelMessageSend(m.ChannelID, "Pardon, même google m'a abandonné.")

		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
	}
}