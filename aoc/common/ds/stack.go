package ds

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func (s *Stack[T]) Empty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() T {
	last := len(s.elements) - 1
	val := s.elements[last]
	s.elements = s.elements[:last]
	return val
}
