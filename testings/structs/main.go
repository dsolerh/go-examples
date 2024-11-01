package main

import "fmt"

type ss struct {
	A int            `json:"a,omitempty"`
	B bool           `json:"b,omitempty"`
	C string         `json:"c,omitempty"`
	D map[string]any `json:"d,omitempty"`
	E []any          `json:"e,omitempty"`
}

func main() {
	s := ss{A: 0, B: false, C: ""}
	// s.D["dddddd"] = true
	s.E = append(s.E, "daniel")
	fmt.Printf("s: %v\n", s)
}
