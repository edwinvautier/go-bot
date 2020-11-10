package test

import (
	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/apis/wit"
	"github.com/edwinvautier/go-bot/commands"
	"testing"
)

func TestCommandBuilder(t *testing.T) {
	type args struct {
		a *wit.Analysis
		s *discordgo.Session
		m *discordgo.MessageCreate
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
			},
			wantErr: true,
		},
		{
			name: "no intent",
			args: args{
				a: &wit.Analysis{
					Intent: []wit.Intent{},
				},
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
