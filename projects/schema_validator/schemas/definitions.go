package schemas

const (
	NumberTypeTag   = "number"
	IntergerTypeTag = "integer"
	StringTypeTag   = "string"
	BooleanTypeTag  = "boolean"
	ObjectTypeTag   = "object"
	ArrayTypeTag    = "array"
)

type NumericTypes interface{ int | float64 }

type LiteralTypes interface{ NumericTypes | string | bool }

type BasicTypeOptions struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

type NumericTypeOptions[T NumericTypes] struct {
	ExclusiveMaximum *T `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum *T `json:"exclusiveMinimum,omitempty"`
	Maximum          *T `json:"maximum,omitempty"`
	Minimum          *T `json:"minimum,omitempty"`
}

type StringContentEncodingOptions string

type StringFormatOptions string

type LiteralTypeSchema struct {
	BasicTypeOptions
	Const any `json:"const"`
}

type BooleanTypeSchema struct {
	BasicTypeOptions
}

type NumericTypeSchema[T NumericTypes] struct {
	BasicTypeOptions
	NumericTypeOptions[T]
}

type IntegerTypeSchema = NumericTypeSchema[int]
type NumberTypeSchema = NumericTypeSchema[float64]

type StringTypeSchema struct {
	BasicTypeOptions
	MaxLength        *int                         `json:"maxLength,omitempty"`        //The maximum string length
	MinLength        *int                         `json:"minLength,omitempty"`        // The minimum string length
	Pattern          string                       `json:"pattern,omitempty"`          // A regular expression pattern this string should match
	Format           StringFormatOptions          `json:"format,omitempty"`           // A format this string should match
	ContentEncoding  StringContentEncodingOptions `json:"contentEncoding,omitempty"`  // The content encoding for this string
	ContentMediaType string                       `json:"contentMediaType,omitempty"` // The content media type for this string
}

type ArrayTypeSchema struct {
	BasicTypeOptions
	Items       any   `json:"items"`
	MinItems    *int  `json:"minItems,omitempty"`    // The maximum number of items in this array
	MaxItems    *int  `json:"maxItems,omitempty"`    // Should this schema contain unique items
	UniqueItems *bool `json:"uniqueItems,omitempty"` // Should this schema contain unique items
}

type ObjectTypeSchema struct {
	BasicTypeOptions
	Properties map[string]any `json:"properties"`
	Required   []string       `json:"required"`
}
