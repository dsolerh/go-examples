package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	encodeMap()
	encodeList()
	encodeArray()
}

type IntConst int

const (
	ONE IntConst = iota + 1
	TWO
	THREE
)

func (i IntConst) String() string {
	switch i {
	case ONE:
		return "ONE"
	case TWO:
		return "TWO"
	case THREE:
		return "THREE"
	default:
		return "INVALID"
	}
}

func (i IntConst) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, i), nil
}

func encodeMap() {
	fmt.Println("encodeMap")
	// json.Marshaler
	var m = make(IntMap[string], 3)
	m[ONE] = "Value 1"
	m[TWO] = "Value 2"
	m[THREE] = "Value 3"

	data, err := json.Marshal(&m)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("data: %s\n", data)
}

func encodeList() {
	fmt.Println("encodeList")
	var list = []IntConst{ONE, TWO, THREE}
	data, err := json.Marshal(&list)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("data: %s\n", data)
}

type IntMap[T any] map[IntConst]T

func (m IntMap[T]) MarshalJSON() ([]byte, error) {
	tm := make(map[string]T, len(m))
	for k, v := range m {
		tm[k.String()] = v
	}
	return json.Marshal(&tm)
}


func encodeArray() {
	fmt.Println("encodeArray")
	type payload struct{
		Array0 [0]int `json:"array_0"`
		Array2 [2]int `json:"array_2"`
		Array4 [4]int `json:"array_4"`
	}
	var value = payload{
		Array0: [0]int{},
		Array2: [2]int{},
		Array4: [4]int{},
	}
	data, err := json.Marshal(&value)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("data: %s\n", data)
}
