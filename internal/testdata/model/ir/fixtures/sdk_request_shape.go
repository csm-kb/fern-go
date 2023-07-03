// This file was auto-generated by Fern from our API Definition.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type SdkRequestShape struct {
	Type            string
	JustRequestBody *HttpRequestBodyReference
	Wrapper         *SdkRequestWrapper
}

func NewSdkRequestShapeFromJustRequestBody(value *HttpRequestBodyReference) *SdkRequestShape {
	return &SdkRequestShape{Type: "justRequestBody", JustRequestBody: value}
}

func NewSdkRequestShapeFromWrapper(value *SdkRequestWrapper) *SdkRequestShape {
	return &SdkRequestShape{Type: "wrapper", Wrapper: value}
}

func (s *SdkRequestShape) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "justRequestBody":
		value := new(HttpRequestBodyReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.JustRequestBody = value
	case "wrapper":
		value := new(SdkRequestWrapper)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Wrapper = value
	}
	return nil
}

func (s SdkRequestShape) MarshalJSON() ([]byte, error) {
	switch s.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", s.Type, s)
	case "justRequestBody":
		var marshaler = struct {
			Type string `json:"type,omitempty"`
			*HttpRequestBodyReference
		}{
			Type:                     s.Type,
			HttpRequestBodyReference: s.JustRequestBody,
		}
		return json.Marshal(marshaler)
	case "wrapper":
		var marshaler = struct {
			Type string `json:"type,omitempty"`
			*SdkRequestWrapper
		}{
			Type:              s.Type,
			SdkRequestWrapper: s.Wrapper,
		}
		return json.Marshal(marshaler)
	}
}

type SdkRequestShapeVisitor interface {
	VisitJustRequestBody(*HttpRequestBodyReference) error
	VisitWrapper(*SdkRequestWrapper) error
}

func (s *SdkRequestShape) Accept(v SdkRequestShapeVisitor) error {
	switch s.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", s.Type, s)
	case "justRequestBody":
		return v.VisitJustRequestBody(s.JustRequestBody)
	case "wrapper":
		return v.VisitWrapper(s.Wrapper)
	}
}
