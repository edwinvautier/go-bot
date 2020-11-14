package commands

import (
	"errors"
	"github.com/edwinvautier/go-bot/apis/youtube"
	log "github.com/sirupsen/logrus"
)

// QueryYoutubeVideoCommand is the struct for our youtube query command
type QueryYoutubeVideoCommand struct {
	gc *GenericCommand
}

// Execute function of the youtube query command
func (command QueryYoutubeVideoCommand) Execute() error {
	if len(command.gc.Analysis.Music) < 1 {
		return errors.New("Missing entity music in analysis")
	}
	music := command.gc.Analysis.Music[0].Value
	if len(music) == 0 {
		return errors.New("No music value specified")
	}

	videos := youtube.SearchByKeywords(music)
	if nil == videos {
		return errors.New("Could not find videos")
	}
	for _, v := range *videos {
		_, err := command.gc.Session.ChannelMessageSend(command.gc.Message.ChannelID, "https://youtu.be/"+v.Id)
		if err != nil {
			log.Error("could not send message to channelId: ", command.gc.Message.ChannelID)
			return errors.New("Could not send message")
		}
	}
	return nil
}
