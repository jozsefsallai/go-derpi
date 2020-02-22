package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SearchImagesResponse contains the full JSON response from Derpibooru's
// images search API
type SearchImagesResponse struct {
	Images []Image `json:"images"`
}

// SearchImages is the equivalent of /api/v1/json/search/images
func (sc SearchClient) SearchImages(query string, page int) ([]Image, error) {
	if page < 1 {
		page = 1
	}

	endpoint := fmt.Sprintf("%s/search/images", sc.Client.BaseURL)
	queryParams := sc.QueryParams()
	queryParams.Add("q", query)

	if page > 1 {
		queryParams.Add("page", strconv.Itoa(page))
	}

	var images SearchImagesResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return []Image{}, err
	}

	json.Unmarshal(response, &images)
	return images.Images, nil
}

// RandomImage will return a random image based on a specific search query
func (sc SearchClient) RandomImage(query string) (Image, error) {
	originalSortField := sc.SortField
	originalPerPage := sc.PerPage

	sc.SortField = "random"
	sc.PerPage = 1

	entries, err := sc.SearchImages(query, 1)

	// reset
	sc.SortField = originalSortField
	sc.PerPage = originalPerPage

	if err != nil {
		return Image{}, err
	}

	if len(entries) == 0 {
		return Image{}, nil
	}

	return entries[0], nil
}
