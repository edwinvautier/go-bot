package connectors

import (
	"github.com/bwmarrin/discordgo"
)

// Discord is our connector to mock "github.com/bwmarrin/discordgo"
type Discord interface {
	ChannelMessageSend(string, string) (*discordgo.Message, error)
}