// This file was auto-generated by Fern from our API Definition.

package user

import (
	context "context"
	fmt "fmt"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/docs/fixtures"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/docs/fixtures/core"
	http "net/http"
	url "net/url"
)

type Client interface {
	GetName(ctx context.Context, userId string, request *fixtures.GetNameRequest) (string, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Returns the username associated with the given userId.
//
// userId uniquely identifies a user.
func (c *client) GetName(ctx context.Context, userId string, request *fixtures.GetNameRequest) (string, error) {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/get-name", userId)

	queryParams := make(url.Values)
	queryParams.Add("filter", fmt.Sprintf("%v", request.Filter))
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := c.header.Clone()
	headers.Add("X-Endpoint-Header", fmt.Sprintf("%v", request.XEndpointHeader))

	var response string
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		headers,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}