package schemas

type StringType struct {
	schema StringTypeSchema
}

func (t *BasicType) String() *StringType {
	var st = new(StringType)
	st.schema.BasicTypeOptions = t.schema
	st.schema.Type = StringTypeTag
	return st
}

func (t *StringType) Schema() any {
	return t.schema
}
