package wordcheck

import (
	"fmt"
)

const (
	runeSize = 4 // 4 bytes
	byteSize = 1 // 1 byte
)

type trieKey interface {
	byte | rune
}

// const Null = 0

// var Substitutions = map[TrieKey][]TrieKey{
// 	// 'a': {'4', '@'},
// }

func newTrie[K trieKey]() *trie[K] {
	return &trie[K]{
		end:    false,
		childs: make(map[K]*trie[K]),
	}
}

var _ fmt.Stringer = (*trie[rune])(nil)
var _ fmt.Stringer = (*trie[byte])(nil)

type trie[K trieKey] struct {
	childs map[K]*trie[K]
	end    bool
}

// Size implements Trie.
// returns the number of bytes stored for the runes in the trie
// this does not represent the actual size in bytes of the trie
// as the map takes it's own space to acount for, but gives an aprox
func (t trie[K]) Size() int64 {
	var size int64

	var unitSize int64
	var k K
	switch any(k).(type) {
	case byte:
		unitSize = byteSize
	case rune:
		unitSize = runeSize
	}

	size += unitSize*int64(len(t.childs)) + 1 // add the byte that takes the bool

	for r := range t.childs {
		size += t.childs[r].Size()
	}

	return size
}

// Rune implements Trie.
func (t *trie[K]) node(k K) *trie[K] {
	return t.childs[k]
}

func (t *trie[K]) put(word []K) {
	node := t
	for _, letter := range word {
		// if the letter is not on the trie add it
		if _, exist := node.childs[letter]; !exist {
			node.childs[letter] = &trie[K]{
				end:    false,
				childs: make(map[K]*trie[K]),
			}
		}

		// add substitution letters
		// for _, subs := range Substitutions[letter] {
		// 	node[subs] = make(trie)
		// }

		// traverse the trie
		node = node.childs[letter]
	}
	// node is going to hold the last child
	node.end = true
}

// Trie implements fmt.Stringer.
func (t *trie[K]) String() string {
	return "implement"
}
