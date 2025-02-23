// This file was auto-generated by Fern from our API Definition.

package dummy

import (
	context "context"
	v2 "github.com/fern-api/stream-go/v2"
	core "github.com/fern-api/stream-go/v2/core"
	http "net/http"
)

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

func (c *Client) GenerateStream(ctx context.Context, request *v2.GenerateStreamRequestzs) (*core.Stream[v2.StreamResponse], error) {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "generate-stream"

	streamer := core.NewStreamer[v2.StreamResponse](c.caller)
	return streamer.Stream(
		ctx,
		&core.StreamParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
			Request: request,
		},
	)
}
