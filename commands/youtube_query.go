package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/apis/youtube"
	log "github.com/sirupsen/logrus"
)

type QueryYoutubeVideo struct {
	analysis *wit.Analysis
}

func (command QueryYoutubeVideo) ExecuteYoutubeSearch(s *discordgo.Session, m *discordgo.MessageCreate) {
	music := command.analysis.Music[0].Value
	if len(music) != 0 {
		log.Info("You asked for music : ", music)
		videos := youtube.SearchByKeywords(music)
		for id, _ := range *videos {
			_, _ = s.ChannelMessageSend(m.ChannelID, "https://youtu.be/"+id)
		}
	}
}
