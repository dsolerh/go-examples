package wordcheck

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type WordChecker interface {
	LoadDictionary(words []string)
	LoadDictionaryFromFile(filename string)
	CheckWord(word string) bool

	Summary() string
}

func New() WordChecker {
	return &checker{
		runeTrie: newTrie[rune](),
		byteTrie: newTrie[byte](),
	}
}

var _ WordChecker = (*checker)(nil)

type checker struct {
	runeTrie *trie[rune]
	byteTrie *trie[byte]
	// metrics
	dictSize     int64
	runeTrieSize int64
	byteTrieSize int64
}

// Summary implements WordChecker.
func (c *checker) Summary() string {
	return fmt.Sprintf(
		"Dictionary size: %d\nByte Trie size: %d\nRune Trie size: %d\nCombined size: %d\n",
		c.dictSize,
		c.byteTrieSize,
		c.runeTrieSize,
		c.byteTrieSize+c.runeTrieSize,
	)
}

// LoadDictionaryFromFile implements WordChecker.
func (c *checker) LoadDictionaryFromFile(filename string) {
	bwordF, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	totalBytes := 0
	scanner := bufio.NewScanner(bwordF)
	for scanner.Scan() {
		word := scanner.Text()
		totalBytes += len(word)
		if isASCII(word) {
			c.byteTrie.put([]byte(word))
		} else {
			c.runeTrie.put([]rune(word))
		}
	}

	// add the sizes
	c.dictSize = int64(totalBytes)
	c.runeTrieSize = c.runeTrie.Size()
	c.byteTrieSize = c.byteTrie.Size()
}

// LoadDictionary implements WordChecker.
func (c *checker) LoadDictionary(words []string) {
	totalBytes := 0
	for _, word := range words {
		totalBytes += len(word)

		if isASCII(word) {
			c.byteTrie.put([]byte(word))
		} else {
			c.runeTrie.put([]rune(word))
		}
	}

	// add the sizes
	c.dictSize = int64(totalBytes)
	c.runeTrieSize = c.runeTrie.Size()
	c.byteTrieSize = c.byteTrie.Size()
}

// CheckWord implements WordChecker.
func (c *checker) CheckWord(word string) bool {
	// check if the word is complete ascii
	if isASCII(word) {
		// do the check with byteTrie
		return contains(c.byteTrie, []byte(word))
	}

	return contains(c.runeTrie, []rune(word))
}

func contains[K trieKey](tr *trie[K], letters []K) bool {
	for i := 0; i < len(letters); i++ {
		t := tr.node(letters[i])

		// if the letter is not in the trie continue
		if t == nil {
			continue
		}

		// if the letter is a terminal letter this means that this letter is
		// considered a bad word, therefore return true
		if t.end {
			return true
		}

		for j := i + 1; j < len(letters); j++ {
			// check if this letter is in the trie
			t = t.node(letters[j])

			// if the letter is not in the trie break
			if t == nil {
				break
			}

			// if the letter is a terminal letter then in the letters there is a
			// chain of letters that form a bad word
			if t.end {
				return true
			}
		}
	}
	return false
}

func isASCII(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
