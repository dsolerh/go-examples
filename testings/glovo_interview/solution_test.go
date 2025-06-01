package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	freqStack := NewStack()

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

func Test_Stack_push(t *testing.T) {
	freqStack := NewStack()

	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(4)
	freqStack.push(5)

	assert.Equal(t, &Stack{
		itemsCounter: 6,
		values: map[int][]int{
			5: {0, 2, 5},
			7: {1, 3},
			4: {4},
		},
	}, freqStack)
}

func Test_Stack_pop(t *testing.T) {
	freqStack := NewStack()

	freqStack.push(5)
	freqStack.push(5)
	freqStack.push(7)
	freqStack.push(3)

	assert.Equal(t, 5, *freqStack.pop())

	freqStack.push(3)
	freqStack.push(5)

	assert.Equal(t, 5, *freqStack.pop())
	assert.Equal(t, 3, *freqStack.pop())
	assert.Equal(t, 3, *freqStack.pop())
	assert.Equal(t, 7, *freqStack.pop())
	assert.Equal(t, 5, *freqStack.pop())
}
