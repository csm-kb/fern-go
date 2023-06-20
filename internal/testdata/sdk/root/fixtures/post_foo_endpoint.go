// Generated by Fern. Do not edit.

package api

import (
	context "context"
	json "encoding/json"
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/root/fixtures/core"
	io "io"
	http "net/http"
)

type postFooEndpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newPostFooEndpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *postFooEndpoint {
	return &postFooEndpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (p *postFooEndpoint) decodeError(statusCode int, body io.Reader) error {
	decoder := json.NewDecoder(body)
	switch statusCode {
	case 409:
		value := new(ConflictError)
		if err := decoder.Decode(value); err != nil {
			return err
		}
		value.StatusCode = statusCode
		return value
	case 422:
		value := new(UnprocessableEntityError)
		if err := decoder.Decode(value); err != nil {
			return err
		}
		value.StatusCode = statusCode
		return value
	}
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (p *postFooEndpoint) Call(ctx context.Context, request *Bar) (*Foo, error) {
	endpointURL := p.url
	response := new(Foo)
	if err := core.DoRequest(
		ctx,
		p.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		p.header,
		p.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}