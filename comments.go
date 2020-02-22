package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// CommentResponse contains the full JSON response from Derpibooru's
// comments API
type CommentResponse struct {
	Comment Comment `json:"comment"`
}

// GetComment is the equivalent of /api/v1/json/comments/:comment_id
func (c Client) GetComment(id int) (Comment, error) {
	endpoint := fmt.Sprintf("%s/comments/%d", c.BaseURL, id)

	var comment CommentResponse
	response, err := apiGet(endpoint, url.Values{})
	if err != nil {
		return Comment{}, err
	}

	json.Unmarshal(response, &comment)
	return comment.Comment, nil
}
