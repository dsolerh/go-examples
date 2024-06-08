package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	type Case struct {
		str   string
		lines bool
		exp   int
	}

	type testCase map[string]Case

	tests := testCase{
		"count words: one line": Case{
			str: "one two three",
			exp: 3,
		},
		"count words: multi lines": Case{
			str: "onw\ndd\nqwer\ndas",
			exp: 4,
		},
		"count lines: one line": Case{
			str:   "one two three",
			lines: true,
			exp:   1,
		},
		"count lines: multi lines": Case{
			str:   "onw\ndd\nqwer\ndas",
			lines: true,
			exp:   4,
		},
	}
	for tname, tcase := range tests {
		t.Run(tname, func(t *testing.T) {
			b := bytes.NewBufferString(tcase.str)
			got := count(b, tcase.lines)

			if got != tcase.exp {
				t.Errorf("got != from expected (%d!=%d)", got, tcase.exp)
			}
		})
	}
}
