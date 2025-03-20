package schemas

import (
	"encoding/json"
	"fmt"
	"schema_validator/utils"
)

type LiteralType struct {
	schema LiteralTypeSchema
}

func (t *BasicType) Literal(value any) *LiteralType {
	var lt = new(LiteralType)
	lt.schema.BasicTypeOptions = t.schema
	lt.schema.Const = value

	switch value.(type) {
	case int, float64:
		lt.schema.Type = NumberTypeTag
	case string:
		lt.schema.Type = StringTypeTag
	case bool:
		lt.schema.Type = BooleanTypeTag
	default:
		panic(fmt.Sprintf("invalid type (%T) for Literal schema", value))
	}

	return lt
}

func (lt *LiteralType) Schema() any {
	return lt.schema
}

func (lt *LiteralType) JSONSchema() string {
	return string(utils.Must(json.Marshal(lt.schema)))
}
