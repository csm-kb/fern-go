// This file was auto-generated by Fern from our API Definition.

package literal

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/fern-api/seed-go/core"
)

type CreateOptionsRequest struct {
	Values map[string]string `json:"values,omitempty"`
}

type GetOptionsRequest struct {
	dryRun bool
}

func (g *GetOptionsRequest) DryRun() bool {
	return g.dryRun
}

func (g *GetOptionsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler GetOptionsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*g = GetOptionsRequest(body)
	g.dryRun = true
	return nil
}

func (g *GetOptionsRequest) MarshalJSON() ([]byte, error) {
	type embed GetOptionsRequest
	var marshaler = struct {
		embed
		DryRun bool `json:"dryRun"`
	}{
		embed:  embed(*g),
		DryRun: true,
	}
	return json.Marshal(marshaler)
}

type GetUndiscriminatedOptionsRequest struct {
	dryRun bool
}

func (g *GetUndiscriminatedOptionsRequest) DryRun() bool {
	return g.dryRun
}

func (g *GetUndiscriminatedOptionsRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler GetUndiscriminatedOptionsRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*g = GetUndiscriminatedOptionsRequest(body)
	g.dryRun = true
	return nil
}

func (g *GetUndiscriminatedOptionsRequest) MarshalJSON() ([]byte, error) {
	type embed GetUndiscriminatedOptionsRequest
	var marshaler = struct {
		embed
		DryRun bool `json:"dryRun"`
	}{
		embed:  embed(*g),
		DryRun: true,
	}
	return json.Marshal(marshaler)
}

type CreateOptionsResponse struct {
	Type    string
	ok      bool
	Options *Options
}

func NewCreateOptionsResponseWithOk() *CreateOptionsResponse {
	return &CreateOptionsResponse{Type: "ok", ok: true}
}

func NewCreateOptionsResponseFromOptions(value *Options) *CreateOptionsResponse {
	return &CreateOptionsResponse{Type: "options", Options: value}
}

func (c *CreateOptionsResponse) Ok() bool {
	return c.ok
}

func (c *CreateOptionsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "ok":
		c.ok = true
	case "options":
		value := new(Options)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Options = value
	}
	return nil
}

func (c CreateOptionsResponse) MarshalJSON() ([]byte, error) {
	switch c.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.Type, c)
	case "ok":
		var marshaler = struct {
			Type string `json:"type"`
			Ok   bool   `json:"value,omitempty"`
		}{
			Type: c.Type,
			Ok:   true,
		}
		return json.Marshal(marshaler)
	case "options":
		var marshaler = struct {
			Type string `json:"type"`
			*Options
		}{
			Type:    c.Type,
			Options: c.Options,
		}
		return json.Marshal(marshaler)
	}
}

type CreateOptionsResponseVisitor interface {
	VisitOk(bool) error
	VisitOptions(*Options) error
}

func (c *CreateOptionsResponse) Accept(visitor CreateOptionsResponseVisitor) error {
	switch c.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", c.Type, c)
	case "ok":
		return visitor.VisitOk(c.ok)
	case "options":
		return visitor.VisitOptions(c.Options)
	}
}

type Options struct {
	Values  map[string]string `json:"values,omitempty"`
	id      string
	enabled bool

	_rawJSON json.RawMessage
}

func (o *Options) Id() string {
	return o.id
}

func (o *Options) Enabled() bool {
	return o.enabled
}

func (o *Options) UnmarshalJSON(data []byte) error {
	type unmarshaler Options
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = Options(value)
	o.id = "options"
	o.enabled = true
	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *Options) MarshalJSON() ([]byte, error) {
	type embed Options
	var marshaler = struct {
		embed
		Id      string `json:"id"`
		Enabled bool   `json:"enabled"`
	}{
		embed:   embed(*o),
		Id:      "options",
		Enabled: true,
	}
	return json.Marshal(marshaler)
}

func (o *Options) String() string {
	if len(o._rawJSON) > 0 {
		if value, err := core.StringifyJSON(o._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

type UndiscriminatedOptions struct {
	typeName        string
	stringLiteral   string
	unknownLiteral  bool
	StringStringMap map[string]string
}

func NewUndiscriminatedOptionsWithStringLiteral() *UndiscriminatedOptions {
	return &UndiscriminatedOptions{typeName: "stringLiteral", stringLiteral: "options"}
}

func NewUndiscriminatedOptionsWithUnknownLiteral() *UndiscriminatedOptions {
	return &UndiscriminatedOptions{typeName: "unknownLiteral", unknownLiteral: true}
}

func NewUndiscriminatedOptionsFromStringStringMap(value map[string]string) *UndiscriminatedOptions {
	return &UndiscriminatedOptions{typeName: "stringStringMap", StringStringMap: value}
}

func (u *UndiscriminatedOptions) StringLiteral() string {
	return u.stringLiteral
}

func (u *UndiscriminatedOptions) UnknownLiteral() bool {
	return u.unknownLiteral
}

func (u *UndiscriminatedOptions) UnmarshalJSON(data []byte) error {
	var valueStringLiteral string
	if err := json.Unmarshal(data, &valueStringLiteral); err == nil {
		if valueStringLiteral == "options" {
			u.typeName = "stringLiteral"
			u.stringLiteral = valueStringLiteral
			return nil
		}
	}
	var valueUnknownLiteral bool
	if err := json.Unmarshal(data, &valueUnknownLiteral); err == nil {
		if valueUnknownLiteral == true {
			u.typeName = "unknownLiteral"
			u.unknownLiteral = valueUnknownLiteral
			return nil
		}
	}
	var valueStringStringMap map[string]string
	if err := json.Unmarshal(data, &valueStringStringMap); err == nil {
		u.typeName = "stringStringMap"
		u.StringStringMap = valueStringStringMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UndiscriminatedOptions) MarshalJSON() ([]byte, error) {
	switch u.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "stringLiteral":
		return json.Marshal("options")
	case "unknownLiteral":
		return json.Marshal(true)
	case "stringStringMap":
		return json.Marshal(u.StringStringMap)
	}
}

type UndiscriminatedOptionsVisitor interface {
	VisitStringLiteral(string) error
	VisitUnknownLiteral(bool) error
	VisitStringStringMap(map[string]string) error
}

func (u *UndiscriminatedOptions) Accept(visitor UndiscriminatedOptionsVisitor) error {
	switch u.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "stringLiteral":
		return visitor.VisitStringLiteral(u.stringLiteral)
	case "unknownLiteral":
		return visitor.VisitUnknownLiteral(u.unknownLiteral)
	case "stringStringMap":
		return visitor.VisitStringStringMap(u.StringStringMap)
	}
}