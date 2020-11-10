package commands

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/apis/youtube"
	log "github.com/sirupsen/logrus"
)

type QueryYoutubeVideo struct {
	analysis *wit.Analysis
	session  *discordgo.Session
	message  *discordgo.MessageCreate
}

func (command QueryYoutubeVideo) Execute() error {
	music := command.analysis.Music[0].Value
	if len(music) != 0 {
		log.Info("You asked for music : ", music)
		videos := youtube.SearchByKeywords(music)
		if nil == videos {
			return errors.New("Could not find videos")
		}
		for id, _ := range *videos {
			_, err := command.session.ChannelMessageSend(command.message.ChannelID, "https://youtu.be/"+id)
			if err != nil {
				return errors.New("Could not send message.")
			}
		}
	}
	return errors.New("No music value specified")
}
