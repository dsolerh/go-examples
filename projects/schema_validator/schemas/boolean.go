package schemas

import (
	"encoding/json"
	"schema_validator/utils"
)

type BooleanType struct {
	schema BooleanTypeSchema
}

func (t *BasicType) Boolean() *BooleanType {
	var bt = new(BooleanType)
	bt.schema.BasicTypeOptions = t.schema
	bt.schema.Type = BooleanTypeTag
	return bt
}

func (t *BooleanType) Schema() any {
	return t.schema
}

func (lt *BooleanType) JSONSchema() string {
	return string(utils.Must(json.Marshal(lt.schema)))
}
