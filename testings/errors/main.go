package main

import (
	"errors"
	"fmt"
)

var _ error = (*ErrCustom)(nil)

type ErrCustom struct {
	code int
	val  string
}

// Error implements error.
func (e *ErrCustom) Error() string {
	return fmt.Sprintf("the code is: %d and the val is: %s", e.code, e.val)
}

func main() {
	err := &ErrCustom{code: 12, val: "super"}

	if customErr := new(*ErrCustom); errors.As(err, customErr) {
		fmt.Printf("customErr: %v\n", *customErr)
	}

	wrappedErr := fmt.Errorf("something custom happened: %w", err)

	if customErr := new(*ErrCustom); errors.As(wrappedErr, customErr) {
		fmt.Printf("customErr: %v\n", *customErr)
	}
}
