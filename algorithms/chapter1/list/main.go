package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)
	intList.PushBack("34")

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Printf("element: %v, type: %T,\n", element.Value, element.Value)
	}
}
