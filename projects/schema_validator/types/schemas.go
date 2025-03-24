package types

type TypeTag string

const (
	NumberTypeTag   TypeTag = "number"
	IntergerTypeTag TypeTag = "integer"
	StringTypeTag   TypeTag = "string"
	BooleanTypeTag  TypeTag = "boolean"
	ObjectTypeTag   TypeTag = "object"
	ArrayTypeTag    TypeTag = "array"
)

type BaseSchema struct {
	Type TypeTag `json:"type"`
}

type LiteralSchema struct {
	BaseSchema
	LiteralOptions
	Const any `json:"const"`
}

type BooleanSchema struct {
	BaseSchema
	BooleanOptions
}

type NumericSchema[T NumericTypes] struct {
	BaseSchema
	NumericOptions[T]
}

type StringSchema struct {
	BaseSchema
	StringOptions
}

type ArraySchema struct {
	BaseSchema
	ArrayOptions
	Items any `json:"items"`
}

type ObjectSchema struct {
	BaseSchema
	ObjectOptions
	Properties map[string]any `json:"properties"`
	Required   []string       `json:"required"`
}
