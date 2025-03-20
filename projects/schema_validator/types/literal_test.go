package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkLiteral[T LiteralTypes](t *testing.T, val T, typ TypeTag) {
	t.Helper()
	t.Run(fmt.Sprintf("Literal[%T](%v) should return a schema of type '%s' with value '%v'", val, val, typ, val), func(t *testing.T) {
		is := assert.New(t)
		s := Literal(val)
		is.Equal(s.schema.Type, typ)
		is.Equal(s.schema.Const, val)
	})
}

func TestLiteral(t *testing.T) {
	checkLiteral(t, 42, IntergerTypeTag)
	checkLiteral(t, 12.3, NumberTypeTag)
	checkLiteral(t, true, BooleanTypeTag)
	checkLiteral(t, "some", StringTypeTag)
}

func TestLiteralWithOptions(t *testing.T) {
	t.Run("should create schema with default options", func(t *testing.T) {
		is := assert.New(t)
		s := Literal(0)
		is.Equal(s.schema.LiteralOptions, LiteralOptions{})
	})
	t.Run("should create schema with passed options", func(t *testing.T) {
		is := assert.New(t)
		op := LiteralOptions{BaseOptions{Description: "A description"}}
		s := Literal(0, &op)
		is.Equal(s.schema.LiteralOptions, op)
	})
}
