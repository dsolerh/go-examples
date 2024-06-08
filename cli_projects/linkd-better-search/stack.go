package main

type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		arr: make([]T, 0),
	}
}

func (s *Stack[T]) Push(el T) {
	s.arr = append(s.arr, el)
}

func (s *Stack[T]) Pop() (T, bool) {
	var last T
	if len(s.arr) == 0 {
		return last, true
	}

	last = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return last, len(s.arr) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}
