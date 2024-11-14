package collections

func Range(start, end int) []int {
	vals := make([]int, end-start)
	for i := 0; start < end; i++ {
		vals[i] = start
		start++
	}
	return vals
}

func Every[T any, S ~[]T](s S, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Some[T any, S ~[]T](s S, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

func Take[T any, S ~[]T](s S, n int) []T {
	if n > len(s) {
		return nil
	}
	return s[:n]
}

func Repeat[T any, S ~[]T](s S, n int) [][]T {
	out := make([][]T, 0, n)
	for ; n > 0; n-- {
		out = append(out, s)
	}
	return out
}

func Join[T any](s [][]T, v T) []T {
	if len(s) == 0 {
		return nil
	}
	out := make([]T, 0, len(s))
	out = append(out, s[0]...)
	if len(s) > 1 {
		for _, val := range s[1:] {
			out = append(out, v)
			out = append(out, val...)
		}
	}
	return out
}

func Flat[T any](s [][]T) []T {
	out := make([]T, 0, len(s))
	for _, val := range s {
		out = append(out, val...)
	}
	return out
}
