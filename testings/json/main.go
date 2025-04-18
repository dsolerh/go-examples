package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// // err := exampleEcodeStream()
	// // if err != nil {
	// // 	fmt.Printf("err: %v\n", err)
	// // }
	// exampleNil()
	// exampleNested()
	// exampleArrayFixed()

	// for i := range 10 {
	// 	fmt.Printf("i: %d i%%2: %d\n", i, i%4)
	// }
	//
	type C uint16
	var c C
	fmt.Printf("c: %016b\n", c)
	c = 1<<8 | 10
	fmt.Printf("c: %016b\n", c)

	fmt.Printf("c: %016b\n", c&0xFF)
	fmt.Printf("c: %016b\n", 0xFF)
}

func exampleArrayFixed() {
	type subSchema struct {
		Name string `json:"name,omitempty"`
	}
	type schema struct {
		Arr4 [4]subSchema `json:"arr,omitempty"`
	}
	val := schema{
		Arr4: [4]subSchema{},
	}
	data, _ := json.Marshal(&val)
	fmt.Printf("data: %s\n", data)
}

func exampleNil() {
	var s any

	b, err := json.Marshal(s)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("b: %s\n", b)
}

func example1() {
	// a := []int{1, 2, 3, 4}
	a := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	data, _ := json.Marshal(&a)
	fmt.Printf("%s\n", data)

	var c any
	err := json.Unmarshal(data, &c)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("c: %v %T\n", c, c)

	switch c.(type) {
	case []int:
		fmt.Println("its an array of ints")
	case []any:
		fmt.Println("its an array")
	default:
		fmt.Println("No idea what ti is")
	}
}

func exampleDecodeStream() error {
	type message struct {
		Name string
		Text string
	}

	f, err := os.Open("json-stream")
	if err != nil {
		return fmt.Errorf("error opening the file: %w", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	msgs := make(chan message, 1)
	errs := make(chan error, 1)

	// Parse the messages concurrently with printing the message.
	go func() {
		defer close(msgs)
		defer close(errs)

		for {
			var m message
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				errs <- err
				return
			}
			msgs <- m
		}
	}()

	for m := range msgs {
		fmt.Printf("message: %+v\n", m)
	}

	if err := <-errs; err != nil {
		return fmt.Errorf("error while streaming: %w", err)
	}

	return nil
}

func exampleEcodeStream() error {
	type message struct {
		Name string
		Text string
	}

	messages := []message{
		{Name: "Ed", Text: "Knock knock."},
		{Name: "Sam", Text: "Who's there?"},
	}

	f, err := os.OpenFile("json-stream", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o0600)
	if err != nil {
		return fmt.Errorf("error while opening the file: %w", err)
	}
	enc := json.NewEncoder(f)

	for _, msg := range messages {
		if err := enc.Encode(msg); err != nil {
			return fmt.Errorf("error while encoding the json: %w", err)
		}
	}

	return nil
}

func exampleNested() {
	type nested struct {
		P1 string `json:"p1"`
	}
	type a struct {
		Nested nested `json:"-"`
	}

	_a := a{Nested: nested{P1: "super"}}

	b, err := json.Marshal(&_a)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("_a: %s\n", b)
}
