package main

import "errors"

// "schema_validator/schemas"

func main() {
	// export const PaymentAmount = function (options?: ObjectOptions) {
	// 	return Type.Object(
	// 		{
	// 			id: Type.Union([Types.PaymentId(), Types.PaymentType()], {
	// 				description: "Id or type of the payment."
	// 			}),
	// 			amount: Type.Optional(
	// 				Type.Integer({
	// 					description: "Amount of the payment."
	// 				})
	// 			),
	// 			params: Type.Optional(
	// 				Type.Record(Type.String(), Type.Any(), {
	// 					description: "Optional additional parameters for the payment.",
	// 					default: {}
	// 				})
	// 			)
	// 		},
	// 		options
	// 	);
	// };
	// export const MatchStakeConfig = makeConfigSchema("matchStake", {
	// 	entryStake: Type.Array(PaymentAmount(), { description: "The initial stake to pay in order to enter a match" }),
	// 	stakeFees: Type.Record(Type.String(), Type.Number({ minimum: 0, maximum: 1 }), {
	// 		description: "For each stake currency, the fee that applies to that currency (default: 0)"
	// 	}),
	// 	maxDouble: Type.Union(
	// 		[
	// 			Type.Literal(1),
	// 			Type.Literal(2),
	// 			Type.Literal(4),
	// 			Type.Literal(8),
	// 			Type.Literal(16),
	// 			Type.Literal(32),
	// 			Type.Literal(64)
	// 		],
	// 		{ description: "The maximum value that the doubling die can reach (inclusive)" }
	// 	)
	// });

}

// func Description(val string) any { return nil }
// func Title(val string) any       { return nil }

// func Minimum[T int | float64](val T) any { return nil }
// func Maximum[T int | float64](val T) any { return nil }

// func Items(val any) any     { return nil }
// func Props(vals ...any) any { return nil }

// func Literal[T int | float64 | string | bool](val T, opts ...any) any { return nil }
// func Integer(opts ...any) any                                         { return nil }
// func Number(opts ...any) any                                          { return nil }
// func String(opts ...any) any                                          { return nil }
// func Boolean(opts ...any) any                                         { return nil }

// func Array(opts ...any) any  { return nil }
// func Object(opts ...any) any { return nil }

// func CustomType(opts ...any) any {
// 	return Object(Props(
// 		"val1", String(),
// 		"val2", Integer(),
// 	), opts)
// }

type LiteralOptions struct{}

func (o *LiteralOptions) Description(s string) *LiteralOptions { return o }

type NumberOptions struct{}

func (o *NumberOptions) Description(s string) *NumberOptions { return o }
func (o *NumberOptions) Minimum(val int) *NumberOptions      { return o }
func (o *NumberOptions) Maximum(val int) *NumberOptions      { return o }

type StringOptions struct{}

func (o *StringOptions) Description(s string) *StringOptions { return o }

type ArrayOptions struct{}

func (o *ArrayOptions) Description(s string) *ArrayOptions { return o }

type ObjectOptions struct{}

func (o *ObjectOptions) Description(s string) *ObjectOptions { return o }

func Literal(val any, opts ...*LiteralOptions) any            { return nil }
func Number(opts ...*NumberOptions) any                       { return nil }
func String(opts ...*StringOptions) any                       { return nil }
func Array(typ any, opts ...*ArrayOptions) any                { return nil }
func Object(props map[string]any, opts ...*ObjectOptions) any { return nil }

type OProp struct {
	key        string
	value      any
	isOptional bool
}

func (p OProp) Optional() OProp { return p }

func Props(props ...OProp) map[string]any { return nil }
func Prop(key string, value any) OProp    { return OProp{key, value, false} }

func CustomObject(opts *ObjectOptions) any { return nil }

func variant3() {
	var _ = Object(
		Props(
			Prop("literal", Literal(12, new(LiteralOptions).Description("just 12"))),
			Prop("number", Number(new(NumberOptions).Description("").Maximum(12).Minimum(2))),
			Prop("maybe", String(new(StringOptions).Description("maybe there"))).Optional(),
			Prop("custom", CustomObject(new(ObjectOptions).Description("custom object"))),
			Prop("array", Array(Number(), new(ArrayOptions).Description(""))),
		),
		new(ObjectOptions).Description(""),
	)
}

type anonymus struct {
	Literal int      `json:"literal"`
	Number  int      `json:"number"`
	Maybe   string   `json:"maybe"`
	Custom  struct{} `json:"custom"`
}

func (o anonymus) validate() error {
	// validates 'literal' value to be '12'
	if o.Literal != 12 {
		return errors.New("value violates literal constraint")
	}

	// validates 'number' value
	if o.Number < 2 || o.Number > 12 {
		return errors.New("value violates number constraint")
	}
	return nil
}

func variant2() {
	// _ = Array(Description(""), Title(""), Items(Integer()))
	// _ = Object(
	// 	Description(""),
	// 	Title(""),
	// 	Props(
	// 		"literal", Literal(12, Description(""), Title("")),
	// 		"integer", Integer(Description(""), Title(""), Minimum(1), Maximum(2)),
	// 		"number", Number(Description(""), Title(""), Minimum(1.2), Maximum(2.4)),
	// 		"string", String(Description(""), Title("")),
	// 		"boolean", Boolean(Description(""), Title("")),
	// 		"array", Array(Description(""), Items(Object(
	// 			Props(
	// 				"prop1", Integer(),
	// 				"prop2", String(),
	// 			),
	// 		))),
	// 		"custom", CustomType(Description(""), Title("")),
	// 	),
	// )
}

func variant1() {
	// var _literal = schemas.Type().Description("a literal value").Literal(12)
	// fmt.Printf("literal: %v\n", _literal)

	// var custom = PaymentAmount("")

	// var _integer = schemas.Integer()
	// fmt.Printf("integer: %v\n", _integer)

	// var _number = schemas.Number()
	// fmt.Printf("number: %v\n", _number)

	// var _boolean = schemas.Boolean()
	// fmt.Printf("boolean: %v\n", _boolean)

	// var _string = schemas.String()
	// fmt.Printf("string: %v\n", _string)

	// var _array = schemas.Array(schemas.Boolean())
	// fmt.Printf("array: %v\n", _array)

	// var _object = schemas.Type().Object().
	// Prop("literal").Of(schemas.Type().Literal(true)).
	// Prop("integer").Of(schemas.Type().Integer().MinValue(1).MaxValue(3)).
	// Prop("number").Optional().Of(schemas.Type().Number().MinValue(0.2).MaxValue(2.0)).
	// Prop("boolean").Of(schemas.Type().Boolean()).
	// Prop("string").Of(schemas.Type().String()).
	// Prop("object").Of(schemas.Type().Object()).
	// Prop("array").Of(schemas.Type().Array().Of(schemas.Type().String()))
	// fmt.Printf("object: %v\n", _object)
}
