// Generated by Fern. Do not edit.

package api

import (
	context "context"
	json "encoding/json"
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/mergent/fixtures/core"
	io "io"
	http "net/http"
)

type postSchedulesEndpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newPostSchedulesEndpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *postSchedulesEndpoint {
	return &postSchedulesEndpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (p *postSchedulesEndpoint) decodeError(statusCode int, body io.Reader) error {
	decoder := json.NewDecoder(body)
	switch statusCode {
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

func (p *postSchedulesEndpoint) Call(ctx context.Context, request *ScheduleNew) (*Schedule, error) {
	endpointURL := p.url
	response := new(Schedule)
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