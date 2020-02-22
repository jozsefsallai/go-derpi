package main

import (
	"net/url"
	"strconv"
)

// SearchClient is a sub-client that can perform searches throguh the
// Derpibooru API
type SearchClient struct {
	Client        Client
	Key           string
	Page          int
	PerPage       int
	SortDirection string
	SortField     string
}

// SearchInitOpts is a structure that contains the options that can be
// used to initialize the SearchClient instance
type SearchInitOpts struct {
	Key           string
	PerPage       int
	SortDirection string
	SortField     string
}

// Search is a method on the Client instance that instantiates a SearchClient
func (c Client) Search(opts SearchInitOpts) SearchClient {
	var searchClient SearchClient
	searchClient.Client = c
	searchClient.PerPage = 25

	if len(c.Key) > 0 {
		searchClient.Key = c.Key
	}

	if len(opts.Key) > 0 {
		searchClient.Key = opts.Key
	}

	if opts.PerPage > 0 {
		searchClient.PerPage = opts.PerPage
	}

	if len(opts.SortDirection) > 0 {
		searchClient.SortDirection = opts.SortDirection
	}

	if len(opts.SortField) > 0 {
		searchClient.SortField = opts.SortField
	}

	return searchClient
}

// QueryParams will return a url.Values instance that contains the default
// query parameters that were set up when calling client.Search()
func (sc SearchClient) QueryParams() url.Values {
	params := url.Values{}

	if sc.PerPage > 1 && sc.PerPage != 25 {
		params.Add("per_page", strconv.Itoa(sc.PerPage))
	}

	if len(sc.Key) > 0 {
		params.Add("key", sc.Key)
	}

	if len(sc.SortDirection) > 0 {
		params.Add("sd", sc.SortDirection)
	}

	if len(sc.SortField) > 0 {
		params.Add("sf", sc.SortField)
	}

	return params
}
