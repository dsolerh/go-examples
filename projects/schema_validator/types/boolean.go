package types

type BooleanOptions struct {
	BaseOptions
}

func (o *BooleanOptions) Description(s string) *BooleanOptions {
	o.BaseOptions.Description = s
	return o
}

type BooleanSchemaBuilder struct {
	schema BooleanSchema
}

func Boolean(opts ...*BooleanOptions) *BooleanSchemaBuilder {
	var bopts BooleanOptions
	if len(opts) != 0 && opts[0] != nil {
		bopts = *opts[0]
	}
	return &BooleanSchemaBuilder{
		schema: BooleanSchema{
			BaseSchema:     BaseSchema{Type: BooleanTypeTag},
			BooleanOptions: bopts,
		},
	}
}

func (sb *BooleanSchemaBuilder) Schema() any {
	return sb.schema
}

func (sb *BooleanSchemaBuilder) ValidationRule(ident string) string {
	return "" // no need for validation rule here, the json parsing will fail if the value is not boolean
}

func (sb *BooleanSchemaBuilder) ValidationError(ident string, val any) string {
	return "" // also no need for this cause there's not going to be an error
}
