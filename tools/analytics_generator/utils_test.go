package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatPascalCase(t *testing.T) {
	tests := []struct {
		word   string
		pascal string
	}{
		{word: "simple", pascal: "Simple"},
		{word: "simpleCompound", pascal: "SimpleCompound"},
		{word: "simpleword", pascal: "Simpleword"},
		{word: "simple_word", pascal: "SimpleWord"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %s", tt.word, tt.pascal), func(t *testing.T) {
			got := formatPascalCase(tt.word)
			assert.Equal(t, tt.pascal, got)
		})
	}
}
func Test_formatCamelCase(t *testing.T) {
	tests := []struct {
		word   string
		pascal string
	}{
		{word: "simple", pascal: "simple"},
		{word: "simpleCompound", pascal: "simpleCompound"},
		{word: "simpleword", pascal: "simpleword"},
		{word: "simple_word", pascal: "simpleWord"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %s", tt.word, tt.pascal), func(t *testing.T) {
			got := formatCamelCase(tt.word)
			assert.Equal(t, tt.pascal, got)
		})
	}
}
