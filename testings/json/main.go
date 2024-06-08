package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// a := []int{1, 2, 3, 4}
	a := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	data, _ := json.Marshal(&a)
	fmt.Printf("%s\n", data)

	var c any
	err := json.Unmarshal(data, &c)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("c: %v %T\n", c, c)

	switch c.(type) {
	case []int:
		fmt.Println("its an array of ints")
	case []any:
		fmt.Println("its an array")
	default:
		fmt.Println("No idea what ti is")
	}

}
