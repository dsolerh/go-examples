package schemas

import (
	"encoding/json"
	"schema_validator/utils"
)

type ObjectType struct {
	schema ObjectTypeSchema
}

func (t *BasicType) Object() *ObjectType {
	var ot = new(ObjectType)
	ot.schema.BasicTypeOptions = t.schema
	ot.schema.Type = ObjectTypeTag
	ot.schema.Properties = make(map[string]any)
	return ot
}

func (t *ObjectType) Schema() any {
	return t.schema
}

func (t *ObjectType) JSONSchema() string {
	return string(utils.Must(json.Marshal(t.schema)))
}

type ObjectTypeProp struct {
	objectType *ObjectType
	isOptional bool
	propName   string
}

func (t *ObjectType) Prop(prop string) *ObjectTypeProp {
	var op = new(ObjectTypeProp)
	op.objectType = t
	op.propName = prop
	return op
}

func (p *ObjectTypeProp) Optional() *ObjectTypeProp {
	p.isOptional = true
	return p
}

func (p *ObjectTypeProp) Of(schema SchemaBuilder) *ObjectType {
	ot := p.objectType
	ot.schema.Properties[p.propName] = schema.Schema()
	if !p.isOptional {
		ot.schema.Required = append(ot.schema.Required, p.propName)
	}
	return ot
}
