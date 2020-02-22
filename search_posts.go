package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchPostsResponse contains the full JSON response from Derpibooru's
// posts search API
type SearchPostsResponse struct {
	Posts []Post `json:"posts"`
}

// SearchPosts is the equivalent of /api/v1/json/search/posts
func (sc SearchClient) SearchPosts(query string, page int) ([]Post, error) {
	if page < 1 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/posts", sc.Client.BaseURL)
	queryParams := sc.QueryParams()
	queryParams.Add("q", query)

	if page > 1 {
		queryParams.Add("page", strconv.Itoa(page))
	}

	var posts SearchPostsResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return []Post{}, err
	}

	json.Unmarshal(response, &posts)
	return posts.Posts, nil
}
