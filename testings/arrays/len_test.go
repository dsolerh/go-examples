package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slice_len(t *testing.T) {
	is := assert.New(t)
	s := make([]int, 0, 5)
	fmt.Printf("s: %v len: %d, cap: %d\n", s, len(s), cap(s))
	is.Equal(0, len(s))
	is.Equal(5, cap(s))

	s = append(s, 1, 2, 3, 4, 5)
	fmt.Printf("s: %v len: %d, cap: %d\n", s, len(s), cap(s))
	is.Equal(5, len(s))
	is.Equal(5, cap(s))

	s = s[:0]
	fmt.Printf("s: %v len: %d, cap: %d\n", s, len(s), cap(s))
	is.Equal(0, len(s))
	is.Equal(5, cap(s))
}

func Test_slice_slice(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[:3]
	assert.Equal(t, a, b)
}
