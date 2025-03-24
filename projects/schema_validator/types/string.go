package types

type StringOptions struct {
	BaseOptions
	MaxLen *int `json:"maxLength,omitempty"` //The maximum string length
	MinLen *int `json:"minLength,omitempty"` // The minimum string length
}

func (o *StringOptions) Description(s string) *StringOptions {
	o.BaseOptions.TypeDescription = s
	return o
}

func (o *StringOptions) MinLength(n int) *StringOptions {
	if n < 0 {
		panic("MinLength should be > 0")
	}
	if o.MaxLen != nil && n >= *o.MaxLen {
		panic("MinLength should be < MaxLength")
	}
	o.MinLen = &n
	return o
}

func (o *StringOptions) MaxLength(n int) *StringOptions {
	if n <= 0 {
		panic("MaxLength should be >= 0")
	}
	if o.MinLen != nil && n <= *o.MinLen {
		panic("MaxLength should be > MinLength")
	}
	o.MinLen = &n
	return o
}

type StringSchemaBuilder struct {
	schema StringSchema
}

func String(opts ...*StringOptions) *StringSchemaBuilder {
	var sopts StringOptions
	if len(opts) != 0 && opts[0] != nil {
		sopts = *opts[0]
	}
	return &StringSchemaBuilder{
		schema: StringSchema{
			BaseSchema:    BaseSchema{Type: StringTypeTag},
			StringOptions: sopts,
		},
	}
}

func (sb *StringSchemaBuilder) Schema() any { return sb.schema }
