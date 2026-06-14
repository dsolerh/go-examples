package main

import "encoding/json"

func Json[T any](str string) (T, error) {
	var value T
	err := json.Unmarshal([]byte(str), &value)
	return value, err
}
