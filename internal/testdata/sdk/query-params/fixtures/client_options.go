// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params/fixtures/core"
)

func ClientWithBaseURL(baseURL string) core.ClientOption {
	return func(opts *core.ClientOptions) {
		opts.BaseURL = baseURL
	}
}

func ClientWithHTTPClient(httpClient core.HTTPClient) core.ClientOption {
	return func(opts *core.ClientOptions) {
		opts.HTTPClient = httpClient
	}
}
