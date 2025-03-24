package types

import "fmt"

type ObjectOptions struct {
	BaseOptions
}

func (o *ObjectOptions) Description(s string) *ObjectOptions {
	o.BaseOptions.TypeDescription = s
	return o
}

type ObjectSchemaBuilder struct {
	schema ObjectSchema
}

type OProp struct {
	key        string
	value      SchemaBuilder
	isOptional bool
}

func Props(props ...OProp) []OProp               { return props }
func Prop(key string, value SchemaBuilder) OProp { return OProp{key: key, value: value} }
func (p OProp) Optional() OProp {
	p.isOptional = true
	return p
}

func Object(props []OProp, opts ...*ObjectOptions) *ObjectSchemaBuilder {
	var oopts ObjectOptions
	if len(opts) != 0 && opts[0] != nil {
		oopts = *opts[0]
	}
	var properties = make(map[string]any, len(props))
	var required = make([]string, 0)
	for _, prop := range props {
		if _, exist := properties[prop.key]; exist {
			panic(fmt.Sprintf("prop %s is already present in the Object schema", prop.key))
		}
		properties[prop.key] = prop.value.Schema()
		if !prop.isOptional {
			required = append(required, prop.key)
		}
	}
	return &ObjectSchemaBuilder{
		schema: ObjectSchema{
			BaseSchema:    BaseSchema{Type: ObjectTypeTag},
			ObjectOptions: oopts,
			Properties:    properties,
			Required:      required,
		},
	}
}

func (sb *ObjectSchemaBuilder) Schema() any { return sb.schema }
