package commands

import (
	"testing"
	"github.com/edwinvautier/go-bot/apis/wit"
)

func TestDispatch(t *testing.T) {
	type args struct {
		a *wit.Analysis
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
			if err := Dispatch(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Dispatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
