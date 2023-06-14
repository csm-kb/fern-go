// Generated by Fern. Do not edit.

package api

import (
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/post-with-path-params/fixtures/core"
	io "io"
)

type UserClient interface{}

type setNameEndpoint struct {
	url    string
	client core.HTTPClient
}

func newsetNameEndpoint(url string, client core.HTTPClient) *setNameEndpoint {
	return &setNameEndpoint{
		url:    url,
		client: client,
	}
}

func (s *setNameEndpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

type setNameV2Endpoint struct {
	url    string
	client core.HTTPClient
}

func newsetNameV2Endpoint(url string, client core.HTTPClient) *setNameV2Endpoint {
	return &setNameV2Endpoint{
		url:    url,
		client: client,
	}
}

func (s *setNameV2Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

type setNameV3Endpoint struct {
	url    string
	client core.HTTPClient
}

func newsetNameV3Endpoint(url string, client core.HTTPClient) *setNameV3Endpoint {
	return &setNameV3Endpoint{
		url:    url,
		client: client,
	}
}

func (s *setNameV3Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

type setNameV4Endpoint struct {
	url    string
	client core.HTTPClient
}

func newsetNameV4Endpoint(url string, client core.HTTPClient) *setNameV4Endpoint {
	return &setNameV4Endpoint{
		url:    url,
		client: client,
	}
}

func (s *setNameV4Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

type setNameV5Endpoint struct {
	url    string
	client core.HTTPClient
}

func newsetNameV5Endpoint(url string, client core.HTTPClient) *setNameV5Endpoint {
	return &setNameV5Endpoint{
		url:    url,
		client: client,
	}
}

func (s *setNameV5Endpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}