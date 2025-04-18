package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mod(t *testing.T) {
	var x = 0
	var prevX = (x + 4 - 1) % 4
	fmt.Printf("prevX: %v\n", prevX)
	assert.Equal(t, 3, prevX)
}
