package schemas

type NoItemArrayType struct {
	schema ArrayTypeSchema
}

type ArrayType struct {
	schema ArrayTypeSchema
}

func (t *BasicType) Array() *NoItemArrayType {
	var at = new(NoItemArrayType)
	at.schema.BasicTypeOptions = t.schema
	at.schema.Type = ArrayTypeTag
	return at
}

func (t *NoItemArrayType) Of(itemsSchema interface{ Schema() any }) *ArrayType {
	var at = new(ArrayType)
	at.schema = t.schema
	at.schema.Items = itemsSchema.Schema()
	return at
}

func (t *ArrayType) Schema() any {
	return t.schema
}
