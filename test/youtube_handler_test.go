package test

import (
	"github.com/edwinvautier/go-bot/apis/youtube"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestSearchByKeywords(t *testing.T) {
	type args struct {
		keyword string
	}

	var videos []youtube.Video
	for i := 0; i < 4; i++ {
		video := youtube.Video{
			Id:    "123",
			Title: "Test title",
		}
		videos = append(videos, video)
	}
	tests := []struct {
		name    string
		args    args
		videos  *[]youtube.Video
		wantErr bool
	}{
		{
			name: "valid query",
			args: args{
				keyword: "keyword test",
			},
			videos:  &videos,
			wantErr: false,
		},
		{
			name: "failing query",
			args: args{
				keyword: "keyword test",
			},
			videos:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _ = youtube.SearchByKeywords(tt.args.keyword); (tt.videos == nil) != tt.wantErr {
				log.Info("Videos:", videos)
				log.Info("Error is required: ", tt.wantErr)
				t.Errorf("SearchByKeywords() wantErr %v", tt.wantErr)
			}
		})
	}
}
