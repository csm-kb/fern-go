// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/fern-api/fern-go/internal/testdata/sdk/basic/fixtures/core"
	http "net/http"
	strings "strings"
)

type Client interface {
	User() UserClient
}

func NewClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) Client {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		header:     options.ToHeader(),
		userClient: NewUserClient(baseURL, httpClient, opts...),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
	userClient UserClient
}

func (c *client) User() UserClient {
	return c.userClient
}
