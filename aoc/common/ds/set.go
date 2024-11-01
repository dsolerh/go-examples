package ds

type empty = struct{}
type Set[T comparable] struct {
	s map[T]empty
}

func NewSet[T comparable, Slice ~[]T](values Slice) *Set[T] {
	set := make(map[T]empty)
	for _, v := range values {
		set[v] = empty{}
	}
	return &Set[T]{
		s: set,
	}
}

func (s *Set[T]) Has(v T) bool {
	_, exist := s.s[v]
	return exist
}
