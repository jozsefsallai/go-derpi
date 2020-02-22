package main

import (
	"net/url"
)

// Client is a Derpibooru API client which contains methods
// that you can use to interact with the Derpibooru API
type Client struct {
	BaseURL string
	Key     string
	Params  url.Values
}

// Init is a method that initializes a client instance
func Init() Client {
	client := Client{}
	client.BaseURL = "https://derpibooru.org/api/v1/json"
	return client
}

// SetURL is a method that allows you to change the base URL
// of the client instance
func (c Client) SetURL(url string) {
	c.BaseURL = url
}

// SetKey will set an API key to the client. This will be ignored
// if some methods have "key" in their query parameters.
func (c Client) SetKey(key string) {
	c.Key = key
}
