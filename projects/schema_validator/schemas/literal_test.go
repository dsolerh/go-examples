package schemas

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralType(t *testing.T) {
	tests := []struct {
		value any
		want  *LiteralType
	}{
		{
			value: 1,
			want: &LiteralType{schema: LiteralTypeSchema{
				BasicTypeOptions: BasicTypeOptions{
					Type: NumberTypeTag,
				},
				Const: 1},
			},
		},
		{
			value: 1.0,
			want: &LiteralType{schema: LiteralTypeSchema{
				BasicTypeOptions: BasicTypeOptions{
					Type: NumberTypeTag,
				},
				Const: 1.0,
			}},
		},
		{
			value: true,
			want: &LiteralType{schema: LiteralTypeSchema{
				BasicTypeOptions: BasicTypeOptions{
					Type: BooleanTypeTag,
				},
				Const: true,
			}},
		},
		{
			value: "some string",
			want: &LiteralType{schema: LiteralTypeSchema{
				BasicTypeOptions: BasicTypeOptions{
					Type: StringTypeTag,
				},
				Const: "some string",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Literal[%T](%v)", tt.value, tt.value), func(t *testing.T) {
			is := assert.New(t)

			got := new(BasicType).Literal(tt.value)

			is.Equal(tt.want, got)
		})
	}
}

func TestLiteralType_JSONSchema(t *testing.T) {
	tests := []struct {
		name            string
		schema          *LiteralType
		wantsJsonSchema string
	}{
		{
			name:            "Literal[bool](true)",
			schema:          new(BasicType).Literal(true),
			wantsJsonSchema: `{"const":true,"type":"boolean"}`,
		},
		{
			name:            "Literal[int](12)",
			schema:          new(BasicType).Literal(12),
			wantsJsonSchema: `{"const":12,"type":"number"}`,
		},
		{
			name:            "Literal[float64](12.3)",
			schema:          new(BasicType).Literal(12.3),
			wantsJsonSchema: `{"const":12.3,"type":"number"}`,
		},
		{
			name:            `Literal[string]("12.3")`,
			schema:          new(BasicType).Literal("12.3"),
			wantsJsonSchema: `{"const":"12.3","type":"string"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := assert.New(t)
			got := tt.schema.JSONSchema()
			is.JSONEq(tt.wantsJsonSchema, got)
		})
	}
}
