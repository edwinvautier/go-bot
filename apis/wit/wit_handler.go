package wit

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	witai "github.com/wit-ai/wit-go"
	"os"
)

// AnalyzeSentence takes a string and uses wit ai API to detect intents inside
func AnalyzeSentence(sentence string) (*Analysis, error) {
	witToken, tokenExist := os.LookupEnv("WIT_TOKEN")
	if !tokenExist {
		return nil, errors.New("Missing environment variable WIT_TOKEN")
	}
	client := witai.NewClient(witToken)

	// Ask the wit API to decode the user request
	msg, err := client.Parse(&witai.MessageRequest{
		Query: sentence,
	})
	if err != nil {
		return nil, err
	}

	// Feed the struct with wit.ai response
	var analysis Analysis
	err = mapstructure.Decode(msg.Entities, &analysis)
	if err != nil {
		return nil, err
	}
	return &analysis, nil
}

// Analysis result
type Analysis struct {
	Intent   []Intent `json:"intent"`
	Location []Entity `json:"location"`
	Music    []Entity `json:"music"`
}

// Intent is the struct we can decode from the wit ai API response to get intents
type Intent struct {
	Confidence float64 `json:"confidence"`
	Value      string  `json:"value"`
}

// Entity is the struct we can decode from the wit ai API response to get entities such as location, music name, etc
type Entity struct {
	Confidence float64 `json:"confidence"`
	Type       string  `json:"type"`
	Value      string  `json:"value"`
}
