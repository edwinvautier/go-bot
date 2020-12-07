package commands

import (
	"errors"
	"strings"
	log "github.com/sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/google"
	"github.com/edwinvautier/go-bot/connectors"
)

// QueryGoogleCommand is the struct for our google query command
type QueryGoogleCommand struct {
	Connector connectors.Discord
	Message   *discordgo.MessageCreate
}

// Execute is the function we use for the google search command
func (command QueryGoogleCommand) Execute() error {
	messageParts := strings.Split(command.Message.Content, "assistant,")
	if len(messageParts) < 2 {
		return errors.New("Unsupported sentence")
	}

	// Call google for our search results
	results, err := google.Search(messageParts[1])
	if err != nil {
		return errors.New("Search failed")
	}

	// Loop inside results and display only the URL inside discord
	_, err = command.Connector.ChannelMessageSend(command.Message.ChannelID, "Voilà ce que j'ai trouvé : ")
	if err != nil {
		log.Error("sendMessageErr: ", err)
	}

	// Send to user the first 3 results
	count := 1
	for _, result := range results {
		if count > 3 {
			break
		}
		_, err = command.Connector.ChannelMessageSend(command.Message.ChannelID, result.URL)
		if err != nil {
			log.Error("sendMessageErr: ", err)
		}
		count++
	}

	return nil
}
