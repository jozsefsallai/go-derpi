package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchGalleriesResponse contains the full JSON response from Derpibooru's
// gallery search API
type SearchGalleriesResponse struct {
	Galleries []Gallery `json:"galleries"`
}

// SearchGalleries is the equivalent of /api/v1/json/search/galleries
func (sc SearchClient) SearchGalleries(query string, page int) ([]Gallery, error) {
	if page < 1 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/galleries", sc.Client.BaseURL)
	queryParams := sc.QueryParams()
	queryParams.Add("q", query)

	if page > 1 {
		queryParams.Add("page", strconv.Itoa(page))
	}

	var galleries SearchGalleriesResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return []Gallery{}, err
	}

	json.Unmarshal(response, &galleries)
	return galleries.Galleries, nil
}
