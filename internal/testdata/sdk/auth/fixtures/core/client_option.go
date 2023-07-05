// This file was auto-generated by Fern from our API Definition.

package core

import (
	fmt "fmt"
	http "net/http"
)

type ClientOption func(*ClientOptions)

type ClientOptions struct {
	BaseURL    string
	HTTPClient HTTPClient
	Custom     string
}

func NewClientOptions() *ClientOptions {
	return &ClientOptions{
		HTTPClient: http.DefaultClient,
	}
}

func (c *ClientOptions) ToHeader() http.Header {
	header := make(http.Header)
	var value string
	if c.Custom != value {
		header.Set("X-API-Key", fmt.Sprintf("%v", c.Custom))
	}
	return header
}
