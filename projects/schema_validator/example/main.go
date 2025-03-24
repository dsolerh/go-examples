package main

import (
	"errors"
	. "schema_validator/types"
)

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

func CustomObject(props ...*ObjectOptions) *ObjectSchemaBuilder {
	return nil
}

func variant3() {
	var _ = Object(
		Props(
			Prop("literal", Literal(12, Options().Description("just 12").Literal())),
			Prop("number", Number(Options().
				Description("").Number().
				Maximum(12).
				Minimum(2))),
			Prop("maybe", String(Options().Description("maybe there").String())).Optional(),
			Prop("custom", CustomObject(Options().Description("custom object").Object())),
			Prop("array", Array(Number(), Options().Description("").Array())),
		),
		Options().Description("").Object(),
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
