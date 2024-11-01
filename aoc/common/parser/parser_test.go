package parser

import (
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {
	data := []byte("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	expected := [][]number{
		{41, 48, 83, 86, 17},
		{83, 86, 6, 31, 17, 9, 48, 53},
	}
	parsed := ParseInput(data,
		Discard(UntilChar(':')),
		Multiples(UntilChar('|'), Discard(Space), Number),
		Multiples(UntilEnd(), Discard(Space), Number),
	)
	if parsed.err != nil {
		t.Errorf("unexpected error in data: %v", parsed.err)
	}

	num0any, _ := parsed.parsed[0].([]any)
	num0 := make([]number, len(num0any))
	for i, n := range num0any {
		num0[i], _ = n.(number)
	}

	if !slices.Equal(expected[0], num0) {
		t.Errorf("expected[0]: %v, got: %v", expected[0], num0)
	}
}
