package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchCommentsResponse contains the full JSON response from Derpibooru's
// comment search API
type SearchCommentsResponse struct {
	Comments []Comment `json:"comments"`
}

// SearchComments is the equivalent of /api/v1/json/search/comments
func (sc SearchClient) SearchComments(query string, page int) ([]Comment, error) {
	if page < 1 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/comments", sc.Client.BaseURL)
	queryParams := sc.QueryParams()
	queryParams.Add("q", query)

	if page > 1 {
		queryParams.Add("page", strconv.Itoa(page))
	}

	var comments SearchCommentsResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return []Comment{}, err
	}

	json.Unmarshal(response, &comments)
	return comments.Comments, nil
}
