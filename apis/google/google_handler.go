package google

import (
    "context"
    "github.com/rocketlaunchr/google-search"
)

// Search function is used to send a query to google
func Search(query string) ([]googlesearch.Result, error) {
	
	ctx := context.Background()
	results, err := googlesearch.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	
	return results, nil
}
