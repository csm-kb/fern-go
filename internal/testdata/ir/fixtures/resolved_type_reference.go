// Generated by Fern. Do not edit.

package ir

import (
	json "encoding/json"
	fmt "fmt"
)

type ResolvedTypeReference struct {
	Type      string
	Container *ContainerType
	Named     *ResolvedNamedType
	Primitive PrimitiveType
	Unknown   any
}

func (r *ResolvedTypeReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	r.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "container":
		var valueUnmarshaler struct {
			Container *ContainerType `json:"container"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		r.Container = valueUnmarshaler.Container
	case "named":
		value := new(ResolvedNamedType)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Named = value
	case "primitive":
		var valueUnmarshaler struct {
			Primitive PrimitiveType `json:"primitive"`
		}
		if err := json.Unmarshal(data, &valueUnmarshaler); err != nil {
			return err
		}
		r.Primitive = valueUnmarshaler.Primitive
	case "unknown":
		value := make(map[string]any)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		r.Unknown = value
	}
	return nil
}

func (r ResolvedTypeReference) MarshalJSON() ([]byte, error) {
	switch r.Type {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", r.Type, r)
	case "container":
		var marshaler = struct {
			Type      string         `json:"_type"`
			Container *ContainerType `json:"container"`
		}{
			Type:      r.Type,
			Container: r.Container,
		}
		return json.Marshal(marshaler)
	case "named":
		var marshaler = struct {
			Type string `json:"_type"`
			*ResolvedNamedType
		}{
			Type:              r.Type,
			ResolvedNamedType: r.Named,
		}
		return json.Marshal(marshaler)
	case "primitive":
		var marshaler = struct {
			Type      string        `json:"_type"`
			Primitive PrimitiveType `json:"primitive"`
		}{
			Type:      r.Type,
			Primitive: r.Primitive,
		}
		return json.Marshal(marshaler)
	case "unknown":
		var marshaler = struct {
			Type    string `json:"_type"`
			Unknown any    `json:"unknown"`
		}{
			Type:    r.Type,
			Unknown: r.Unknown,
		}
		return json.Marshal(marshaler)
	}
}

type ResolvedTypeReferenceVisitor interface {
	VisitContainer(*ContainerType) error
	VisitNamed(*ResolvedNamedType) error
	VisitPrimitive(PrimitiveType) error
	VisitUnknown(any) error
}

func (r *ResolvedTypeReference) Accept(v ResolvedTypeReferenceVisitor) error {
	switch r.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", r.Type, r)
	case "container":
		return v.VisitContainer(r.Container)
	case "named":
		return v.VisitNamed(r.Named)
	case "primitive":
		return v.VisitPrimitive(r.Primitive)
	case "unknown":
		return v.VisitUnknown(r.Unknown)
	}
}