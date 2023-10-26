// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	metricsclient "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/organization/metrics/client"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header

	Metrics *metricsclient.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
		Metrics:    metricsclient.NewClient(opts...),
	}
}

func (c *Client) Check(ctx context.Context, id string) (*fixtures.Organization, error) {
	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"organization/%v", id)

	var response *fixtures.Organization
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return nil, err
	}
	return response, nil
}