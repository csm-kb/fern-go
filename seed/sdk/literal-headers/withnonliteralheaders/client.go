// This file was auto-generated by Fern from our API Definition.

package withnonliteralheaders

import (
	context "context"
	fmt "fmt"
	seedgo "github.com/fern-api/seed-go"
	core "github.com/fern-api/seed-go/core"
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

func (c *Client) Get(ctx context.Context, request *seedgo.WithNonLiteralHeadersRequest) error {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "with-non-literal-headers"

	headers := c.header.Clone()
	headers.Add("nonLiteralEndpointHeader", fmt.Sprintf("%v", request.NonLiteralEndpointHeader))
	headers.Add("literalEndpointHeader", fmt.Sprintf("%v", "endpoint header"))
	headers.Add("trueEndpointHeader", fmt.Sprintf("%v", true))

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: headers,
		},
	); err != nil {
		return err
	}
	return nil
}
