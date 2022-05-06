package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := helloResponse{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := helloResponse{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkHelloHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := helloResponse{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}
