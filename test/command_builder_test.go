package test

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/commands"
	"github.com/edwinvautier/go-bot/connectors"
	"testing"
)

type discordSessionMock struct {}

func (session *discordSessionMock) ChannelMessageSend(channelID string, message string) (*discordgo.Message, error) {
	return nil, nil
}
func TestCommandBuilder(t *testing.T) {
	type args struct {
		a *wit.Analysis
		s connectors.Discord
		m *discordgo.MessageCreate
	}

	discordMock := discordSessionMock{}
	discordMessage := discordgo.Message {
		ChannelID: "1",
	}
	discordMessageCreate := discordgo.MessageCreate {
		&discordMessage,
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "intent music",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{{
						Value: "listen",
					}},
				},
				s: &discordMock,
				m: &discordMessageCreate,
			},
			wantErr: false,
		},
		{
			name: "intent meteo",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{{
						Value: "meteo",
					}},
				},
				s: &discordMock,
				m: &discordMessageCreate,
			},
			wantErr: false,
		},
		{
			name: "unknown intent",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{{
						Value: "ukulele",
					}},
				},
				s: &discordMock,
				m: &discordMessageCreate,
			},
			wantErr: false,
		},
		{
			name: "no intent",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{},
				},
				s: &discordMock,
				m: &discordMessageCreate,
			},
			wantErr: true,
		},
		{
			name: "empty intent",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{{
						Value: "",
					}},
				},
				s: &discordMock,
				m: &discordMessageCreate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := commands.Build(tt.args.a, tt.args.s, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
