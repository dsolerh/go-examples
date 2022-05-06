package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

var buff bytes.Buffer

func init() {
	buff.WriteString(`{"name":"Daniel"}`)
}

func BenchmarkBodyVariable(b *testing.B) {
	b.ResetTimer()

	var request helloRequest

	for i := 0; i < b.N; i++ {
		json.Unmarshal(buff.Bytes(), &request)
	}
}

func BenchmarkBodyDecoder(b *testing.B) {
	b.ResetTimer()

	var request helloRequest

	for i := 0; i < b.N; i++ {
		decoder := json.NewDecoder(&buff)
		decoder.Decode(request)
	}
}

func BenchmarkBodyDecoderReference(b *testing.B) {
	b.ResetTimer()

	var request helloRequest

	for i := 0; i < b.N; i++ {
		decoder := json.NewDecoder(&buff)
		decoder.Decode(&request)
	}
}
