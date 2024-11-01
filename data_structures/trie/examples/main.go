package main

import (
	"bufio"
	"data_structures/trie"
	"fmt"
	"os"
)

var badWordsTrie trie.Trie

func init() {
	badWordsTrie = trie.NewTrie()

	bwordF, err := os.Open("bad-words.txt")
	if err != nil {
		panic(err)
	}
	totalBytes := 0
	scanner := bufio.NewScanner(bwordF)
	for scanner.Scan() {
		word := scanner.Text()
		totalBytes += len(word)
		badWordsTrie.Put(word)
	}
	fmt.Printf("totalBytes: %v\n", totalBytes)
}

func containsBadWords(letters []rune) bool {
	for i := 0; i < len(letters); i++ {
		t := badWordsTrie.Node(letters[i])

		// if the letter is not in the trie continue
		if t.Empty() {
			continue
		}

		// if the letter is a terminal letter this means that this letter is
		// considered a bad word, therefore return true
		if t.Last() {
			return true
		}

		for j := i + 1; j < len(letters); j++ {
			// check if this letter is in the trie
			t = t.Node(letters[j])

			// if the letter is not in the trie break
			if t.Empty() {
				break
			}

			// if the letter is a terminal letter then in the letters there is a
			// chain of letters that form a bad word
			if t.Last() || t.Node(trie.Null).Last() {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Printf("badWordsTrie.Size(): %v bytes \n", badWordsTrie.Size())
	bWords := []string{
		"fuck",
		"f4ck",
	}
	for _, word := range bWords {
		ok := containsBadWords([]rune(word))
		fmt.Printf("'%s' is a bad word (%t)\n", word, ok)
	}
}
