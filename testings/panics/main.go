package main

import (
	"fmt"
)

func main() {
	catcher()
	fmt.Println("after recovered")

	fmt.Printf("willReturn(): %v\n", willReturn())
}

// Panic panics with a divide by zero
func willPanic() {
	var zero int // default value is 0
	// zero, err := strconv.ParseInt("0", 10, 64)
	// if err != nil {
	// 	panic(err)
	// }
	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher calls Panic
func catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	func() {
		func() {
			willPanic()
		}()
	}()
}

func willReturn() (ret string) {
	defer superRecover(&ret)
	willPanic()
	return "nothing happened"
}

func superRecover(s *string) {
	if r := recover(); r != nil {
		*s = "panic on the disco"
	}
}
