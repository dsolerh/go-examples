package types

import "fmt"

type LiteralOptions struct {
	BaseOptions
}

func (o *LiteralOptions) Description(s string) *LiteralOptions {
	o.BaseOptions.Description = s
	return o
}

type LiteralSchemaBuilder struct {
	schema LiteralSchema
}

func Literal[T LiteralTypes](val T, opts ...*LiteralOptions) *LiteralSchemaBuilder {
	var lopts LiteralOptions
	if len(opts) != 0 && opts[0] != nil {
		lopts = *opts[0]
	}
	var typ TypeTag
	switch any(val).(type) {
	case int:
		typ = IntergerTypeTag
	case float64:
		typ = NumberTypeTag
	case bool:
		typ = BooleanTypeTag
	case string:
		typ = StringTypeTag
	default:
		panic(fmt.Sprintf("the value %T(%v) does not respect the types for 'LiteralTypes'", val, val))
	}
	return &LiteralSchemaBuilder{
		schema: LiteralSchema{
			BaseSchema:     BaseSchema{Type: typ},
			LiteralOptions: lopts,
			Const:          val,
		},
	}
}

func (sb *LiteralSchemaBuilder) Schema() any {
	return sb.schema
}

func (sb *LiteralSchemaBuilder) ValidationRule(ident string) string {
	// returns an expresion that if evaluated in golang will return true
	// if the value passed is invalid
	return fmt.Sprintf(`%s != %v`, ident, sb.schema.Const)
}

func (sb *LiteralSchemaBuilder) ValidationError(ident string, val any) string {
	return fmt.Sprintf(`invalid literal value '%v' for %s, expected '%v'`, val, ident, sb.schema.Const)
}
