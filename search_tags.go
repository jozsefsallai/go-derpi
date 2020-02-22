package derpi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchTagsResponse contains the full JSON response from Derpibooru's
// tag search API
type SearchTagsResponse struct {
	Tags []Tag `json:"tags"`
}

// SearchTags is the equivalent of /api/v1/json/search/tags
func (sc SearchClient) SearchTags(query string, page int) ([]Tag, error) {
	if page < 1 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/tags", sc.Client.BaseURL)
	queryParams := sc.QueryParams()
	queryParams.Add("q", query)

	if page > 1 {
		queryParams.Add("page", strconv.Itoa(page))
	}

	var tags SearchTagsResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return []Tag{}, err
	}

	json.Unmarshal(response, &tags)
	return tags.Tags, nil
}
