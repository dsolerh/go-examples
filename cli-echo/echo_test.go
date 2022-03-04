package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	testCases := []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two"}, "one\ttwo\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"}, // NOTE: wrong expectation!
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("echo(%v, %q, %q)", tC.newline, tC.sep, tC.args)
		t.Run(desc, func(t *testing.T) {
			out = new(bytes.Buffer) // capture output
			if err := echo(tC.newline, tC.sep, tC.args); err != nil {
				t.Errorf("%s failed: %v", desc, err)
				return
			}
			got := out.(*bytes.Buffer).String()
			if got != tC.want {
				t.Errorf("got %q, want %q", got, tC.want)
			}
		})
	}
}
