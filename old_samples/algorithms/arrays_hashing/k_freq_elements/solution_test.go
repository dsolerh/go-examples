package kfreqelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKElements(t *testing.T) {
	is := assert.New(t)
	nums := []int{1, 2, 2, 2, 2, 3, 3, 3, 3}
	k := 2
	want := []int{2, 3}
	got := TopKElements(nums, k)
	is.ElementsMatch(want, got)
}
