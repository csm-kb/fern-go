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

type setNameV3Endpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newSetNameV3Endpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *setNameV3Endpoint {
	return &setNameV3Endpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (s *setNameV3Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (s *setNameV3Endpoint) Call(ctx context.Context, userId string, request *SetNameRequestV3) (*SetNameRequestV3Body, error) {
	endpointURL := fmt.Sprintf(s.url, userId)
	response := new(SetNameRequestV3Body)
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