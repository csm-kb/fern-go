// This file was auto-generated by Fern from our API Definition.

package core

import (
	http "net/http"
)

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	BaseURL    string
	HTTPClient HTTPClient
	Bearer     string
}

func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		HTTPClient: http.DefaultClient,
	}
}

func (c *ClientOptions) ToHeader() http.Header {
	header := make(http.Header)
	if c.Bearer != "" {
		header.Set("Authorization", "Bearer "+c.Bearer)
	}
	return header
}
