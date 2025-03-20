package utils

import "fmt"

func Must[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
	return val
}
