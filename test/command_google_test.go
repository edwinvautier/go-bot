package test

import (
	"github.com/edwinvautier/go-bot/commands"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/edwinvautier/go-bot/connectors"
)

func TestQueryGoogleCommand_Execute(t *testing.T) {
	type fields struct {
		Connector connectors.Discord
		Message   *discordgo.MessageCreate
	}
	
	// Declare params
	discordMock := discordSessionMock{}
	discordMessage := discordgo.Message {
		ChannelID: "1",
		Content: "assistant, comment fabriquer un sous-marin ?",
	}
	discordMessageCreate := discordgo.MessageCreate {
		&discordMessage,
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Empty query",
			fields: fields{
				Connector: &discordMock,
				Message: &discordgo.MessageCreate {
					&discordgo.Message {
						ChannelID: "1",
						Content: "",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "true query",
			fields: fields{
				Connector: &discordMock,
				Message: &discordMessageCreate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			command := commands.QueryGoogleCommand{
				Connector: tt.fields.Connector,
				Message:   tt.fields.Message,
			}
			if err := command.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("QueryGoogleCommand.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
