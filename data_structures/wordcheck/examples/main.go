package main

import (
	"data_structures/wordcheck"
	"fmt"
)

func main() {
	wchk := wordcheck.New()
	wchk.LoadDictionaryFromFile("bad-words.txt")
	fmt.Printf("wchk.Summary(): %s\n", wchk.Summary())
	bWords := []string{
		"fuck",
		"f4ck",
	}
	for _, word := range bWords {
		ok := wchk.CheckWord(word)
		fmt.Printf("'%s' is a bad word (%t)\n", word, ok)
	}
}
