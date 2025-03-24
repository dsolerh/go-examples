package types

type BaseOptions struct {
	TypeDescription string `json:"description,omitempty"`
}

func Options() *BaseOptions { return new(BaseOptions) }

func (o *BaseOptions) Description(s string) *BaseOptions {
	o.TypeDescription = s
	return o
}

func (o *BaseOptions) Literal() *LiteralOptions {
	lo := new(LiteralOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) Number() *NumberOptions {
	lo := new(NumberOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) Integer() *IntegerOptions {
	lo := new(IntegerOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) Boolean() *BooleanOptions {
	lo := new(BooleanOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) String() *StringOptions {
	lo := new(StringOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) Array() *ArrayOptions {
	lo := new(ArrayOptions)
	lo.BaseOptions = *o
	return lo
}

func (o *BaseOptions) Object() *ObjectOptions {
	lo := new(ObjectOptions)
	lo.BaseOptions = *o
	return lo
}
