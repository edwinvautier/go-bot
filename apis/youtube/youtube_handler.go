package youtube

import (
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"os"

	"google.golang.org/api/youtube/v3"
)

var (
	YService   *youtube.Service
	maxResults = flag.Int64("max-results", 5, "Max YouTube results")
)

func ClientInit() {
	youtubeToken, exist := os.LookupEnv("YOUTUBE_TOKEN")
	if !exist {
		log.Fatal("Missing environment variable : YOUTUBE_TOKEN")
	}
	flag.Parse()

	ctx := context.Background()

	opt := option.WithAPIKey(youtubeToken)

	service, err := youtube.NewService(ctx, opt)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	YService = service
}

func SearchByKeywords(query string) *map[string]string {
	var parts []string
	parts = append(parts, "id")
	parts = append(parts, "snippet")
	// Make the API call to YouTube.

	//query := flag.String("query", "Google", "Search term")
	fmt.Println("Query", query)
	call := YService.Search.List(parts).
		Q(query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		return nil
	}
	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	return &videos
}
