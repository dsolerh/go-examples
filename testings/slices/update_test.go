package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_update_slice(t *testing.T) {
	var st = struct{ arr []int }{
		arr: []int{1, 2, 3, 4},
	}
	var fn = func(arr []int) (found bool) {
		for i := range arr {
			if arr[i] == 2 {
				found = true
				arr[i] = 99
				return
			}
		}
		return
	}

	fn(st.arr)
	assert.Equal(t, []int{1, 99, 3, 4}, st.arr)
}
