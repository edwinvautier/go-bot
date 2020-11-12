package youtube

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"os"

	"google.golang.org/api/youtube/v3"
)

var (
	maxResults = flag.Int64("max-results", 5, "Max YouTube results")
	Yservice   *youtube.Service
)

// ClientInit initialize a youtube service.
// Assign Yservice global variable to fetched service
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

	Yservice = service
}

// SearchByKeyword let you search a youtube video from keywords.
// Takes a keyword as argument.
// Return a pointer of an Video array.
func SearchByKeywords(keyword string) *[]Video {
	var parts []string
	parts = append(parts, "id")
	parts = append(parts, "snippet")
	// Make the API call to YouTube.

	//query := flag.String("query", "Google", "Search term")
	call := Yservice.Search.List(parts).
		Q(keyword).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		return nil
	}
	var videos []Video
	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			video := Video{
				Id:    item.Id.VideoId,
				Title: item.Snippet.Title,
			}
			videos = append(videos, video)
		}
	}

	return &videos
}

// Video object to reflect the data gathered from the api.
type Video struct {
	Id    string
	Title string
}
