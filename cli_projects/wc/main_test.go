package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	type Case struct {
		str string
		exp int
	}

	type testCase map[string]Case

	tests := testCase{
		"one line": Case{
			str: "one two three",
			exp: 3,
		},
		"multi lines": Case{
			str: "onw\ndd\nqwer\ndas",
			exp: 4,
		},
	}
	for tname, tcase := range tests {
		t.Run(tname, func(t *testing.T) {
			b := bytes.NewBufferString(tcase.str)
			got := count(b)

			if got != tcase.exp {
				t.Errorf("got != from expected (%d!=%d)", got, tcase.exp)
			}
		})
	}
}
