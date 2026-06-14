package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{"a":"Value 1","b":"Value 2"}`)
	type AA struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	type Wrap struct {
		a *AA
	}
	var val Wrap = Wrap{a: new(AA)}

	err := json.Unmarshal(data, val.a)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("val: %v\n", val.a)
}
