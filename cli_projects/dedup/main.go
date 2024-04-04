package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type empty struct{}

type NoDupWords map[string]empty

func NewNoDupWords() NoDupWords {
	return make(NoDupWords)
}

func (w NoDupWords) Add(word string) {
	w[word] = empty{}
}

func (w NoDupWords) All() []string {
	allWords := make([]string, 0, len(w))
	for word := range w {
		allWords = append(allWords, word)
	}
	return allWords
}

func main() {
	filename := flag.String("f", "", "put a filename")
	flag.Parse()

	file, err := os.Open(*filename)
	check(err)

	data, err := io.ReadAll(file)
	check(err)

	words := NewNoDupWords()
	for _, word := range strings.Split(string(data), "\n") {
		words.Add(word)
	}

	fmt.Printf("words: %#v\n", words.All())
}
