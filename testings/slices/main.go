package main

import (
	"fmt"
	"slices"
)

func main() {
	type S struct {
		v1 byte
		v2 byte
	}
	arr := []S{{1, 2}, {2, 3}, {3, 4}}
	fmt.Printf("arr: %v\n", arr)
	arr2 := slices.DeleteFunc(arr, func(e S) bool { return e == S{2, 3} })
	fmt.Printf("arr2: %v\n", arr2)
	// re define the array cause the call to DeleteFunc will actually modify the array values if it finds one to delete
	arr = []S{{1, 2}, {2, 3}, {3, 4}}
	arr3 := slices.DeleteFunc(arr, func(e S) bool { return e == S{1, 4} })
	fmt.Printf("arr3: %v\n", arr3)
}
