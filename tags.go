package derpi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TagResponse contains the full JSON response from Derpibooru's
// tags API
type TagResponse struct {
	Tag Tag `json:"tag"`
}

// TagQueryParams contains the allowed query parameters for
// Derpibooru's tags API
type TagQueryParams struct {
	Key string
}

func getTag(client Client, id string, params TagQueryParams) (Tag, error) {
	endpoint := fmt.Sprintf("%s/tags/%s", client.BaseURL, id)
	queryParams := url.Values{}

	if len(client.Key) > 0 {
		queryParams.Add("key", client.Key)
	}

	if len(params.Key) > 0 {
		queryParams.Add("key", params.Key)
	}

	var tag TagResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return Tag{}, err
	}

	json.Unmarshal(response, &tag)
	return tag.Tag, nil
}

// GetTag is the equivalent of /api/v1/json/tags/:tag_id
func (c Client) GetTag(id string) (Tag, error) {
	return getTag(c, id, TagQueryParams{})
}

// GetTagOpts is the equivalent of /api/v1/json/tags/:tag_id
// with query parameters
func (c Client) GetTagOpts(id string, params TagQueryParams) (Tag, error) {
	return getTag(c, id, params)
}
