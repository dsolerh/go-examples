package types

type ArrayOptions struct {
	BaseOptions
}

func (o *ArrayOptions) Description(s string) *ArrayOptions {
	o.BaseOptions.TypeDescription = s
	return o
}

type ArraySchemaBuilder struct {
	schema ArraySchema
}

func Array(typ SchemaBuilder, opts ...*ArrayOptions) *ArraySchemaBuilder {
	var aopts ArrayOptions
	if len(opts) != 0 && opts[0] != nil {
		aopts = *opts[0]
	}
	return &ArraySchemaBuilder{
		schema: ArraySchema{
			BaseSchema:   BaseSchema{Type: ArrayTypeTag},
			ArrayOptions: aopts,
			Items:        typ.Schema(),
		},
	}
}

func (sb *ArraySchemaBuilder) Schema() any { return sb.schema }
