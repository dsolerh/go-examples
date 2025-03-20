package schemas

type BasicType struct {
	schema BasicTypeOptions
}

func Type() *BasicType {
	return new(BasicType)
}

func (t *BasicType) Description(description string) *BasicType {
	t.schema.Description = description
	return t
}
