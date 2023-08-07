// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
)

type ConflictError struct {
	*core.APIError
	Body *Error
}

func (c *ConflictError) UnmarshalJSON(data []byte) error {
	var body *Error
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.StatusCode = 409
	c.Body = body
	return nil
}

func (c *ConflictError) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

type UnprocessableEntityError struct {
	*core.APIError
	Body *Error
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *Error
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}