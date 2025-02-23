// This file was auto-generated by Fern from our API Definition.

package service

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	seedgo "github.com/fern-api/seed-go"
	core "github.com/fern-api/seed-go/core"
	io "io"
	multipart "mime/multipart"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

func (c *Client) Post(ctx context.Context, file io.Reader, maybeFile io.Reader, request *seedgo.MyRequest) error {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL

	requestBuffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(requestBuffer)
	fileFilename := "file_filename"
	if named, ok := file.(interface{ Name() string }); ok {
		fileFilename = named.Name()
	}
	filePart, err := writer.CreateFormFile("file", fileFilename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(filePart, file); err != nil {
		return err
	}
	if maybeFile != nil {
		maybeFileFilename := "maybeFile_filename"
		if named, ok := maybeFile.(interface{ Name() string }); ok {
			maybeFileFilename = named.Name()
		}
		maybeFilePart, err := writer.CreateFormFile("maybeFile", maybeFileFilename)
		if err != nil {
			return err
		}
		if _, err := io.Copy(maybeFilePart, maybeFile); err != nil {
			return err
		}
	}
	if request.MaybeString != nil {
		if err := writer.WriteField("maybeString", fmt.Sprintf("%v", *request.MaybeString)); err != nil {
			return err
		}
	}
	if err := writer.WriteField("integer", fmt.Sprintf("%v", request.Integer)); err != nil {
		return err
	}
	if request.MaybeInteger != nil {
		if err := writer.WriteField("maybeInteger", fmt.Sprintf("%v", *request.MaybeInteger)); err != nil {
			return err
		}
	}
	if err := core.WriteMultipartJSON(writer, "listOfStrings", request.ListOfStrings); err != nil {
		return err
	}
	if err := core.WriteMultipartJSON(writer, "setOfStrings", request.SetOfStrings); err != nil {
		return err
	}
	if request.OptionalListOfStrings != nil {
		if err := core.WriteMultipartJSON(writer, "optionalListOfStrings", request.OptionalListOfStrings); err != nil {
			return err
		}
	}
	if request.OptionalSetOfStrings != nil {
		if err := core.WriteMultipartJSON(writer, "optionalSetOfStrings", request.OptionalSetOfStrings); err != nil {
			return err
		}
	}
	if err := core.WriteMultipartJSON(writer, "maybeList", request.MaybeList); err != nil {
		return err
	}
	if request.OptionalMaybeList != nil {
		if err := core.WriteMultipartJSON(writer, "optionalMaybeList", request.OptionalMaybeList); err != nil {
			return err
		}
	}
	if err := core.WriteMultipartJSON(writer, "maybeListOrSet", request.MaybeListOrSet); err != nil {
		return err
	}
	if request.OptionalMaybeListOrSet != nil {
		if err := core.WriteMultipartJSON(writer, "optionalMaybeListOrSet", request.OptionalMaybeListOrSet); err != nil {
			return err
		}
	}
	if err := core.WriteMultipartJSON(writer, "listOfObjects", request.ListOfObjects); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	c.header.Set("Content-Type", writer.FormDataContentType())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
			Request: requestBuffer,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) JustFile(ctx context.Context, file io.Reader) error {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "just-file"

	requestBuffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(requestBuffer)
	fileFilename := "file_filename"
	if named, ok := file.(interface{ Name() string }); ok {
		fileFilename = named.Name()
	}
	filePart, err := writer.CreateFormFile("file", fileFilename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(filePart, file); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	c.header.Set("Content-Type", writer.FormDataContentType())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:     endpointURL,
			Method:  http.MethodPost,
			Headers: c.header,
			Request: requestBuffer,
		},
	); err != nil {
		return err
	}
	return nil
}
