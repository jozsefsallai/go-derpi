package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// ImageResponse contains the full JSON response from Derpibooru's
// images API
type ImageResponse struct {
	Image Image `json:"image"`
}

// ImageQueryParams contains the allowed query parameters for
// Derpibooru's images API
type ImageQueryParams struct {
	Key      string
	FilterID int
}

func getImage(client Client, id int, params ImageQueryParams) (Image, error) {
	endpoint := fmt.Sprintf("%s/images/%d", client.BaseURL, id)
	queryParams := url.Values{}

	if len(client.Key) > 0 {
		queryParams.Add("key", client.Key)
	}

	if len(params.Key) > 0 {
		queryParams.Add("key", params.Key)
	}

	if params.FilterID > 0 {
		queryParams.Add("filter_id", strconv.Itoa(params.FilterID))
	}

	var image ImageResponse
	response, err := apiGet(endpoint, queryParams)
	if err != nil {
		return Image{}, err
	}

	json.Unmarshal(response, &image)
	return image.Image, nil
}

// GetImage is the equivalent of /api/v1/json/images/:image_id
func (c Client) GetImage(id int) (Image, error) {
	return getImage(c, id, ImageQueryParams{})
}

// GetImageOpts is the equivalent of /api/v1/json/images/:image_id
// with query parameters
func (c Client) GetImageOpts(id int, params ImageQueryParams) (Image, error) {
	return getImage(c, id, params)
}
