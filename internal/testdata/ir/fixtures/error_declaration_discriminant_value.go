// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ErrorDeclarationDiscriminantValue struct {
	Type       string
	Property   *NameAndWireValue
	StatusCode any
}

func (e *ErrorDeclarationDiscriminantValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "property":
		value := new(NameAndWireValue)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Property = value
	case "statusCode":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.StatusCode = value
	}
	return nil
}

func (e ErrorDeclarationDiscriminantValue) MarshalJSON() ([]byte, error) {
	switch e.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "property":
		var marshaler = struct {
			Type string `json:"type"`
			*NameAndWireValue
		}{
			Type:             e.Type,
			NameAndWireValue: e.Property,
		}
		return json.Marshal(marshaler)
	case "statusCode":
		var marshaler = struct {
			Type       string `json:"type"`
			StatusCode any    `json:"statusCode"`
		}{
			Type:       e.Type,
			StatusCode: e.StatusCode,
		}
		return json.Marshal(marshaler)
	}
}

type ErrorDeclarationDiscriminantValueVisitor interface {
	VisitProperty(*NameAndWireValue) error
	VisitStatusCode(any) error
}

func (e *ErrorDeclarationDiscriminantValue) Accept(v ErrorDeclarationDiscriminantValueVisitor) error {
	switch e.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", e.Type, e)
	case "property":
		return v.VisitProperty(e.Property)
	case "statusCode":
		return v.VisitStatusCode(e.StatusCode)
	}
}