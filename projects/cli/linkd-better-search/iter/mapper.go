package iter

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func (iter Iter[T]) Map(mapFn func(T) T) Iter[T] {
	return Map(iter, mapFn)
}

// Map returns a new iterator that yields the results of applying the provided
// function to the input iterator.
func Map[T, U any](iter Iter[T], mapFn func(T) U) Iter[U] {
	return func() (U, bool) {
		next, ok := iter()
		if ok {
			return mapFn(next), true
		}

		return Default[U](), false
	}
}
