// This file was auto-generated by Fern from our API Definition.

package ir

type HttpPath struct {
	Head  string          `json:"head,omitempty"`
	Parts []*HttpPathPart `json:"parts,omitempty"`
}
