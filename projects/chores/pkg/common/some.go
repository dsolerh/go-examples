package common

func Some[T any, Slice ~[]T](s Slice, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}
