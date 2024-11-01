package trie

import (
	"encoding/json"
	"fmt"
)

const (
	runeSize = 4 // 4 bytes
)

type TrieKey = rune

const Null TrieKey = 0

var Substitutions = map[TrieKey][]TrieKey{
	// 'a': {'4', '@'},
}

type Trie interface {
	Put(string)
	Has(string) bool
	Node(TrieKey) Node
	Size() int64
}

type Node interface {
	Node(TrieKey) Node
	Empty() bool
	Last() bool
}

func NewTrie() Trie {
	return make(trie)
}

var _ fmt.Stringer = trie(nil)
var _ Trie = trie(nil)
var _ Node = trie(nil)

type trie map[rune]trie

// Size implements Trie.
// returns the number of bytes stored for the runes in the trie
// this does not represent the actual size in bytes of the trie
// as the map takes it's own space to acount for, but gives an aprox
func (t trie) Size() int64 {
	var size int64

	size += runeSize * int64(len(t))

	for r := range t {
		size += t[r].Size()
	}

	return size
}

// Empty implements TrieNode.
func (t trie) Empty() bool {
	return t == nil
}

// Last implements TrieNode.
func (t trie) Last() bool {
	return t != nil && len(t) == 0
}

// Rune implements Trie.
func (t trie) Node(r TrieKey) Node {
	return t[r]
}

// Has implements Trie.
func (t trie) Has(word string) bool {
	node := t
	var letter rune
	for _, letter = range word {
		// if the letter is not present in the map return false
		if _, exist := node[letter]; !exist {
			return false
		}
		node = node[letter]
	}
	return letter == Null || node[Null] != nil
}

// Put implements Trie.
func (t trie) Put(word string) {
	node := t
	for _, letter := range word {
		// if the letter is not on the trie add it
		if _, exist := node[letter]; !exist {
			node[letter] = make(trie)
		}

		// add substitution letters
		for _, subs := range Substitutions[letter] {
			node[subs] = make(trie)
		}

		// traverse the trie
		node = node[letter]
	}
	// node is going to hold the last child
	node[Null] = trie{}
}

// Trie implements fmt.Stringer.
func (t trie) String() string {
	b, err := json.MarshalIndent(&t, "", "\t")
	if err != nil {
		return fmt.Sprintf("could not represent to string: %v", err)
	}
	return string(b)
}
