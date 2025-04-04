package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo[T any] struct {
	Bar
	Baz *T
}

type Bar struct {
	Prop1 int
	prop2 string
}

type Baz struct {
	Prop1 bool
}

func init() {
	gob.Register(foo[Baz]{})
}

func Test_gob_encode(t *testing.T) {
	state := map[string]any{
		"foo": foo[Baz]{
			Bar: Bar{
				Prop1: 42,
				prop2: "prop2",
			},
			Baz: &Baz{
				Prop1: true,
			},
		},
	}
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(state)
	assert.NoError(t, err)

	fmt.Printf("buf.String(): %v\n", buf.String())

	var decodedState map[string]any
	var expectedState = map[string]any{
		"foo": foo[Baz]{
			Bar: Bar{
				Prop1: 42,
			},
			Baz: &Baz{
				Prop1: true,
			},
		},
	}
	err = gob.NewDecoder(buf).Decode(&decodedState)
	assert.NoError(t, err)
	assert.Equal(t, expectedState, decodedState)
}
