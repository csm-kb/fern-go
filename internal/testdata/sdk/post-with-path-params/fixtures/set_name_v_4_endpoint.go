// Generated by Fern. Do not edit.

package api

import (
	context "context"
	errors "errors"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/post-with-path-params/fixtures/core"
	io "io"
	http "net/http"
)

type setNameV4Endpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newSetNameV4Endpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *setNameV4Endpoint {
	return &setNameV4Endpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (s *setNameV4Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (s *setNameV4Endpoint) Call(ctx context.Context, userId string, request *SetNameRequestV4) (string, error) {
	endpointURL := fmt.Sprintf(s.url, userId)
	var response string
	if err := core.DoRequest(
		ctx,
		s.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		s.header,
		s.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}