package client

import (
	"net/http"

	"github.com/fern-api/fern-go/internal/generator/sdk/core"
)

// ---
// This file is not used as a template for code generation. It is only included so that
// the client_test.go template can actually run as a real test file in this repository.
// ---

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}
