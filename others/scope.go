package main

import "fmt"

func scope_1() {
	var list []func()
	for i := 0; i < 3; i++ {
		i := i
		list = append(list, func() { fmt.Printf("i: %v\n", i) })
	}
	for _, v := range list {
		v()
	}
}
