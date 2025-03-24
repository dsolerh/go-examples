package types

type NumericOptions[T NumericTypes] struct {
	BaseOptions
	MinVal     *T `json:"minimum,omitempty"`
	MaxVal     *T `json:"maximum,omitempty"`
	ExclMinVal *T `json:"exclusiveMinimum,omitempty"`
	ExclMaxVal *T `json:"exclusiveMaximum,omitempty"`
}

type NumberOptions = NumericOptions[float64]
type IntegerOptions = NumericOptions[int]

func (o *NumericOptions[T]) Description(s string) *NumericOptions[T] {
	o.BaseOptions.TypeDescription = s
	return o
}

func (o *NumericOptions[T]) Minimum(val T) *NumericOptions[T] {
	if o.ExclMinVal != nil {
		panic("should not set Minimum and ExclusiveMinimum at the same time")
	}
	o.MinVal = &val
	return o
}

func (o *NumericOptions[T]) Maximum(val T) *NumericOptions[T] {
	if o.ExclMaxVal != nil {
		panic("should not set Maximum and ExclusiveMaximum at the same time")
	}
	o.MaxVal = &val
	return o
}

func (o *NumericOptions[T]) ExclusiveMinimum(val T) *NumericOptions[T] {
	if o.MinVal != nil {
		panic("should not set ExclusiveMinimum and Minimum at the same time")
	}
	o.ExclMinVal = &val
	return o
}

func (o *NumericOptions[T]) ExclusiveMaximum(val T) *NumericOptions[T] {
	if o.MaxVal != nil {
		panic("should not set ExclusiveMaximum and Maximum at the same time")
	}
	o.ExclMaxVal = &val
	return o
}

type NumericSchemaBuilder[T NumericTypes] struct {
	schema NumericSchema[T]
}

// func Boolean(opts ...*BooleanOptions) *BooleanSchemaBuilder {
func Integer(opts ...*IntegerOptions) *NumericSchemaBuilder[int] {
	var nopts IntegerOptions
	if len(opts) != 0 && opts[0] != nil {
		nopts = *opts[0]
	}
	return &NumericSchemaBuilder[int]{
		schema: NumericSchema[int]{
			BaseSchema:     BaseSchema{Type: IntergerTypeTag},
			NumericOptions: nopts,
		},
	}
}

func Number(opts ...*NumberOptions) *NumericSchemaBuilder[float64] {
	var nopts NumberOptions
	if len(opts) != 0 && opts[0] != nil {
		nopts = *opts[0]
	}
	return &NumericSchemaBuilder[float64]{
		schema: NumericSchema[float64]{
			BaseSchema:     BaseSchema{Type: NumberTypeTag},
			NumericOptions: nopts,
		},
	}
}

func (sb *NumericSchemaBuilder[T]) Schema() any { return sb.schema }
