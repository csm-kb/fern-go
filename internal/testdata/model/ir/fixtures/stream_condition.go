// This file was auto-generated by Fern from our API Definition.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type StreamCondition struct {
	Type string
	// The name of a boolean query parameter. If it is true, the response
	// should be streamed. Otherwise, it should not be streamed.
	QueryParameterKey string
	// The name of a boolean property on the request. If it is true, the response
	// should be streamed. Otherwise, it should not be streamed.
	RequestPropertyKey string
}

func NewStreamConditionFromQueryParameterKey(value string) *StreamCondition {
	return &StreamCondition{Type: "queryParameterKey", QueryParameterKey: value}
}

func NewStreamConditionFromRequestPropertyKey(value string) *StreamCondition {
	return &StreamCondition{Type: "requestPropertyKey", RequestPropertyKey: value}
}

func (s *StreamCondition) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "queryParameterKey":
		var valueUnmarshaler struct {
			QueryParameterKey string `json:"value,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		s.QueryParameterKey = valueUnmarshaler.QueryParameterKey
	case "requestPropertyKey":
		var valueUnmarshaler struct {
			RequestPropertyKey string `json:"value,omitempty"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		s.RequestPropertyKey = valueUnmarshaler.RequestPropertyKey
	}
	return nil
}

func (s StreamCondition) MarshalJSON() ([]byte, error) {
	switch s.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", s.Type, s)
	case "queryParameterKey":
		var marshaler = struct {
			Type              string `json:"type,omitempty"`
			QueryParameterKey string `json:"value,omitempty"`
		}{
			Type:              s.Type,
			QueryParameterKey: s.QueryParameterKey,
		}
		return json.Marshal(marshaler)
	case "requestPropertyKey":
		var marshaler = struct {
			Type               string `json:"type,omitempty"`
			RequestPropertyKey string `json:"value,omitempty"`
		}{
			Type:               s.Type,
			RequestPropertyKey: s.RequestPropertyKey,
		}
		return json.Marshal(marshaler)
	}
}

type StreamConditionVisitor interface {
	VisitQueryParameterKey(string) error
	VisitRequestPropertyKey(string) error
}

func (s *StreamCondition) Accept(v StreamConditionVisitor) error {
	switch s.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", s.Type, s)
	case "queryParameterKey":
		return v.VisitQueryParameterKey(s.QueryParameterKey)
	case "requestPropertyKey":
		return v.VisitRequestPropertyKey(s.RequestPropertyKey)
	}
}
