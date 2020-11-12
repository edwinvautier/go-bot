package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/apis/youtube"
	log "github.com/sirupsen/logrus"
)

type QueryYoutubeVideoCommand struct {
	analysis *wit.Analysis
	session  *discordgo.Session
	message  *discordgo.MessageCreate
}

func (command QueryYoutubeVideoCommand) Execute() error {
	music := command.analysis.Music[0].Value
	if len(music) == 0 {
		return errors.New("No music value specified")
	}

	log.Info("You asked for music : ", music)
	videos := youtube.SearchByKeywords(music)
	if nil == videos {
		return errors.New("Could not find videos")
	}
	for _, v := range *videos {
		_, err := command.session.ChannelMessageSend(command.message.ChannelID, "https://youtu.be/"+v.Id)
		if err != nil {
			log.Error("could not send message to channelId: ", command.message.ChannelID)
			return errors.New("Could not send message.")
		}
	}
	return nil
}
