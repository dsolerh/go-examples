package main

// Filter filters the slice s based on the fn provided. if
// fn returns true the value is kept in the filtered slice
func Filter[T any](s []T, fn func(T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range s {
		if fn(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
