package api

import (
	json "encoding/json"
	fmt "fmt"
	bar "github.com/fern-api/fern-go/internal/testdata/packages/generated/go/bar"
)

type Union struct {
	Type         string
	Value        *Value
	AnotherValue *Value
	Bar          *bar.Bar
	AnotherBar   *bar.Bar
}

func (x *Union) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "value":
		value := new(Value)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Value = value
	case "anotherValue":
		var valueUnmarshaler struct {
			AnotherValue *Value `json:"anotherValue"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.AnotherValue = valueUnmarshaler.AnotherValue
	case "bar":
		value := new(bar.Bar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Bar = value
	case "anotherBar":
		var valueUnmarshaler struct {
			AnotherBar *bar.Bar `json:"anotherBar"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		x.AnotherBar = valueUnmarshaler.AnotherBar
	}
	return nil
}

func (x Union) MarshalJSON() ([]byte, error) {
	switch x.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "value":
		var marshaler = struct {
			Type string `json:"type"`
			*Value
		}{
			Type:  x.Type,
			Value: x.Value,
		}
		return json.Marshal(marshaler)
	case "anotherValue":
		var marshaler = struct {
			Type         string `json:"type"`
			AnotherValue *Value `json:"anotherValue"`
		}{
			Type:         x.Type,
			AnotherValue: x.AnotherValue,
		}
		return json.Marshal(marshaler)
	case "bar":
		var marshaler = struct {
			Type string `json:"type"`
			*bar.Bar
		}{
			Type: x.Type,
			Bar:  x.Bar,
		}
		return json.Marshal(marshaler)
	case "anotherBar":
		var marshaler = struct {
			Type       string   `json:"type"`
			AnotherBar *bar.Bar `json:"anotherBar"`
		}{
			Type:       x.Type,
			AnotherBar: x.AnotherBar,
		}
		return json.Marshal(marshaler)
	}
}

type UnionVisitor interface {
	VisitValue(*Value) error
	VisitAnotherValue(*Value) error
	VisitBar(*bar.Bar) error
	VisitAnotherBar(*bar.Bar) error
}

func (x *Union) Accept(v UnionVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "value":
		return v.VisitValue(x.Value)
	case "anotherValue":
		return v.VisitAnotherValue(x.AnotherValue)
	case "bar":
		return v.VisitBar(x.Bar)
	case "anotherBar":
		return v.VisitAnotherBar(x.AnotherBar)
	}
}