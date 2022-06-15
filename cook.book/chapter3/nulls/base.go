package nulls

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name":"Aaron"}`
	fullJsonBlob = `{"name":"Aaron","age":0}`
)

// Example is a basic struct with age and name fields
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

// BaseEncoding shows encoding and
// decoding with normal types
func BaseEncoding() error {
	e := Example{}
	// note that no age = 0 age
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshall, no age: %+v\n", e)
	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)
	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}
