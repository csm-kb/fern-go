// This file was auto-generated by Fern from our API Definition.

package api

import (
	context "context"
	base64 "encoding/base64"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params-complex/fixtures/core"
	uuid "github.com/gofrs/uuid/v5"
	http "net/http"
	url "net/url"
	time "time"
)

type UserClient interface {
	GetUsername(ctx context.Context, request *GetUsersRequest) (*User, error)
}

func NewUserClient(opts ...core.ClientOption) UserClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &userClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type userClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (u *userClient) GetUsername(ctx context.Context, request *GetUsersRequest) (*User, error) {
	baseURL := ""
	if u.baseURL != "" {
		baseURL = u.baseURL
	}
	endpointURL := baseURL + "/" + "/user"

	queryParams := make(url.Values)
	var idDefaultValue uuid.UUID
	if request.Id != idDefaultValue {
		queryParams.Add("id", fmt.Sprintf("%v", request.Id))
	}
	var dateDefaultValue time.Time
	if request.Date != dateDefaultValue {
		queryParams.Add("date", fmt.Sprintf("%v", request.Date.Format("2006-01-02")))
	}
	var deadlineDefaultValue time.Time
	if request.Deadline != deadlineDefaultValue {
		queryParams.Add("deadline", fmt.Sprintf("%v", request.Deadline.Format(time.RFC3339)))
	}
	if request.Bytes != nil {
		queryParams.Add("bytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(request.Bytes)))
	}
	var optionalIdDefaultValue *uuid.UUID
	if request.OptionalId != optionalIdDefaultValue {
		queryParams.Add("optionalId", fmt.Sprintf("%v", *request.OptionalId))
	}
	var optionalDateDefaultValue *time.Time
	if request.OptionalDate != optionalDateDefaultValue {
		queryParams.Add("optionalDate", fmt.Sprintf("%v", request.OptionalDate.Format("2006-01-02")))
	}
	var optionalDeadlineDefaultValue *time.Time
	if request.OptionalDeadline != optionalDeadlineDefaultValue {
		queryParams.Add("optionalDeadline", fmt.Sprintf("%v", request.OptionalDeadline.Format(time.RFC3339)))
	}
	var optionalBytesDefaultValue *[]byte
	if request.OptionalBytes != optionalBytesDefaultValue {
		queryParams.Add("optionalBytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(*request.OptionalBytes)))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	response := new(User)
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		u.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}