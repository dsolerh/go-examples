package main

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

func main() {
	aa := []byte("WyIiLDEsMzBd")
	fmt.Printf("aa: %v\n", aa)
	valid := utf8.Valid(aa)
	fmt.Printf("valid: %v\n", valid)
	var da = []any{}
	err := json.Unmarshal(aa, &da)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("da: %v\n", da)
}
