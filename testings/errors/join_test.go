package main

import (
	"errors"
	"fmt"
)

func Example_errors_join() {
	var errA = errors.New("error A happen")
	var errB = fmt.Errorf("error B happen with %d", 42)

	fmt.Printf("joined: %s", errors.Join(errA, errB).Error())
	// Output: joined: error A happen
	// error B happen with 42
}
