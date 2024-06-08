package iter

type Iter[T any] func() (T, bool)

func (iter Iter[T]) Collect() []T {
	res := make([]T, 0)
	for next, ok := iter(); ok; next, ok = iter() {
		// next, ok := iter()

		// if !ok {
		// 	break
		// }

		res = append(res, next)
	}
	return res
}

// Elems returns an iterator over the values of the provided slice.
func IntoIter[T any](seq []T) Iter[T] {
	index := -1
	return func() (T, bool) {
		index++
		if index < len(seq) {
			return seq[index], true
		}
		return Default[T](), false
	}
}
