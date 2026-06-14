package main

type SType string

const (
	TObject  SType = "object"
	TString  SType = "string"
	TInteger SType = "integer"
	TBoolean SType = "boolean"
)

var typesMap = map[SType]string{
	TString:  "string",
	TBoolean: "bool",
	TInteger: "int",
}

type TProperties = map[string]SchemaType

type SchemaType struct {
	Type       SType       `json:"type"`
	Properties TProperties `json:"properties"`
	Const      string      `json:"const"`
	// Required   []any       `json:"required"` // TODO
}

type SelfInfo struct {
	Namespace string `json:"namespace,omitempty"`
	Category  string `json:"category,omitempty"`
	Name      string `json:"name,omitempty"`
	Version   string `json:"version,omitempty"`
}

type MetaInfo struct {
	Vertical bool   `json:"vertical,omitempty"`
	Dynamic  string `json:"dynamic,omitempty"`
}

type SchemaMeta struct {
	Self   SelfInfo `json:"self,omitempty"`
	Meta   MetaInfo `json:"meta,omitempty"`
	Schema string   `json:"$schema,omitempty"`
	Id     string   `json:"$id,omitempty"`
}

type Schema struct {
	SchemaType
	SchemaMeta
}
