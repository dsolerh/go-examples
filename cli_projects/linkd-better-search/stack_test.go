package main

import (
	"reflect"
	"testing"
)

func Test_Stack_Push(t *testing.T) {
	stack := NewStack[any]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expectArr := []any{1, 2, 3}
	if !reflect.DeepEqual(stack.arr, expectArr) {
		t.Errorf("expected: %v, but got: %v", expectArr, stack.arr)
	}
}

func Test_Stack_Pop(t *testing.T) {
	stack := &Stack[any]{
		arr: []any{1, 2},
	}

	expectPop1 := 2
	if pop1, _ := stack.Pop(); expectPop1 != pop1 {
		t.Errorf("expected: %v, but got: %v", expectPop1, pop1)
	}

	expectPop2 := 1
	pop2, last := stack.Pop()
	if expectPop2 != pop2 {
		t.Errorf("expected: %v, but got: %v", expectPop2, pop2)
	}
	if !last {
		t.Errorf("Expected last element")
	}
}
