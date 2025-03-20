package schemas

type NumericType[T NumericTypes] struct {
	schema NumericTypeSchema[T]
}

func (t *BasicType) Integer() *NumericType[int] {
	var nt = new(NumericType[int])
	nt.schema.BasicTypeOptions = t.schema
	nt.schema.Type = IntergerTypeTag
	return nt
}

func (t *BasicType) Number() *NumericType[float64] {
	var nt = new(NumericType[float64])
	nt.schema.BasicTypeOptions = t.schema
	nt.schema.Type = NumberTypeTag
	return nt
}

func (t *NumericType[int]) Schema() any {
	return t.schema
}

func (t *NumericType[T]) MinValue(val T) *NumericType[T] {
	t.schema.Minimum = &val
	return t
}

func (t *NumericType[T]) MaxValue(val T) *NumericType[T] {
	t.schema.Maximum = &val
	return t
}
