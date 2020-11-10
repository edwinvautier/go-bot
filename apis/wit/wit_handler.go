package wit

import (
	witai "github.com/wit-ai/wit-go"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/mitchellh/mapstructure"
)

func AnalyzeSentence(sentence string) Analysis {
	witToken, tokenExist := os.LookupEnv("WIT_TOKEN")
	if !tokenExist {
		log.Error("Missing environment variable WIT_TOKEN")
		return Analysis{}
	}
	client := witai.NewClient(witToken)

	msg, err := client.Parse(&witai.MessageRequest{
		Query: sentence,
	})

	if err != nil {
		log.Error("Error while parsing request: ", err)
		return Analysis{}
	}
	var analysis Analysis
	mapstructure.Decode(msg.Entities, &analysis)

	return analysis
}

type Analysis struct {
	Intent 		[]Intent	`json:"intent"`
	Location	[]Entity	`json:"location"`
	Music		[]Entity	`json:"music"`
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