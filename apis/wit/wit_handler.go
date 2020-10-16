package wit

import (
	"encoding/json"
	witai "github.com/wit-ai/wit-go"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/mitchellh/mapstructure"
)

func AnalyzeSentence(sentence string) {
	witToken, tokenExist := os.LookupEnv("WIT_TOKEN")
	if !tokenExist {
		log.Error("Missing environment variable WIT_TOKEN")
		return
	}
	client := witai.NewClient(witToken)

	msg, err := client.Parse(&witai.MessageRequest{
		Query: sentence,
	})

	if err != nil {
		log.Error("Error while parsing request: ", err)
		return
	}

	var analysis Analysis
	data, _ := json.MarshalIndent(msg.Entities, "", " ")
	log.Info(string(data))
	mapstructure.Decode(msg.Entities, &analysis)
	log.Info(analysis)
}

type Analysis struct {
	Intents 	[]Entity	`json:"intent"`
	Location	[]Entity	`json:"location"`
}

type Intent struct {
	Confidence 	float64		`json:"confidence"`
	Value		string		`json:"value"`
}

type Entity struct {
	Confidence	float64		`json:"confidence"`
	Type		bool		`json:"type"`
	Value		string		`json:"value"`
}