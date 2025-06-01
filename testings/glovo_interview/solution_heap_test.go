package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StackWithHeap(t *testing.T) {
	freqStack := NewStackWithHeap()

	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(4)
	freqStack.push(5)

	assert.Equal(t, 5, *freqStack.pop())
	assert.Equal(t, 7, *freqStack.pop())
	assert.Equal(t, 5, *freqStack.pop())
	assert.Equal(t, 4, *freqStack.pop())
	assert.Equal(t, 7, *freqStack.pop())
	assert.Equal(t, 5, *freqStack.pop())
	assert.Nil(t, freqStack.pop())
}

func Test_StackWithHeap_push(t *testing.T) {
	freqStack := NewStackWithHeap()

	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(4)
	freqStack.push(5)

	assert.Equal(t, &StackWithHeap{
		itemsCounter: 6,
		values: map[int]int{
			5: 3,
			7: 2,
			4: 1,
		},
		heap: &maxHeap{
			{val: 5, count: 1, order: 0},
			{val: 7, count: 1, order: 1},
			{val: 4, count: 1, order: 4},
			{val: 5, count: 2, order: 2},
			{val: 7, count: 2, order: 3},
			{val: 5, count: 3, order: 5},
		},
	}, freqStack)
}
