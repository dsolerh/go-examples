package parser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseNumber(t *testing.T) {
	tests := []struct {
		data []byte
		want number
	}{
		{[]byte("0"), 0},
		{[]byte("01"), 1},
		{[]byte("012"), 12},
		{[]byte("1"), 1},
		{[]byte("12"), 12},
		{[]byte("123"), 123},
		{[]byte("1234"), 1234},
		{[]byte("12345"), 12345},
		{[]byte("123456"), 123456},
		{[]byte("1234567"), 1234567},
		{[]byte("12345678"), 12345678},
		{[]byte("123456789"), 123456789},
		{[]byte("1234567890"), 1234567890},
		{[]byte("12345678901"), 12345678901},
		{[]byte("123456789012"), 123456789012},
		{[]byte("1234567890123"), 1234567890123},
		{[]byte("12345678901234"), 12345678901234},
		{[]byte("123456789012345"), 123456789012345},
		{[]byte("1234567890123456"), 1234567890123456},
		{[]byte("12345678901234567"), 12345678901234567},
		{[]byte("123456789012345678"), 123456789012345678},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("str: %s -> num: %d", tt.data, tt.want), func(t *testing.T) {
			if got := ParseNumber(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
