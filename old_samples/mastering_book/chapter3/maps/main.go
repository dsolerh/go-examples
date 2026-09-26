package main

import (
	"fmt"
)

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "value"

	// range
	for k, v := range aMap {
		fmt.Println("key: ", k, " value: ", v)
	}
}
