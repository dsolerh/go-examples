package wordcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trie_put(t *testing.T) {
	for _, key := range []any{rune(0), byte(0)} {
		if _, ok := key.(byte); ok {
			t.Run("running for trie of type byte", triePut[byte])
		} else {
			t.Run("running for trie of type rune", triePut[rune])
		}
	}
}

func triePut[K trieKey](t *testing.T) {
	t.Helper()
	tests := []struct {
		tr     *trie[K]
		wantTr *trie[K]
		name   string
		word   []K
	}{
		{
			name: "should add the string to an empty trie",
			word: []K{K('f'), K('o'), K('o')}, // []K("foo")
			tr: &trie[K]{
				end:    false,
				childs: make(map[K]*trie[K]),
			},
			wantTr: &trie[K]{
				end: false,
				childs: map[K]*trie[K]{
					'f': {
						end: false,
						childs: map[K]*trie[K]{
							'o': {
								end: false,
								childs: map[K]*trie[K]{
									'o': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "should add the string to an existing trie",
			word: []K{K('f'), K('o'), K('z')}, // []K("foz")
			tr: &trie[K]{
				end: false,
				childs: map[K]*trie[K]{
					'f': {
						end: false,
						childs: map[K]*trie[K]{
							'o': {
								end: false,
								childs: map[K]*trie[K]{
									'o': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
								},
							},
						},
					},
				},
			},
			wantTr: &trie[K]{
				end: false,
				childs: map[K]*trie[K]{
					'f': {
						end: false,
						childs: map[K]*trie[K]{
							'o': {
								end: false,
								childs: map[K]*trie[K]{
									'o': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
									'z': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "should mark the end of the string in the trie without adding it",
			word: []K{K('f'), K('o')}, // []K("fo")
			tr: &trie[K]{
				end: false,
				childs: map[K]*trie[K]{
					'f': {
						end: false,
						childs: map[K]*trie[K]{
							'o': {
								end: false,
								childs: map[K]*trie[K]{
									'o': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
								},
							},
						},
					},
				},
			},
			wantTr: &trie[K]{
				end: false,
				childs: map[K]*trie[K]{
					'f': {
						end: false,
						childs: map[K]*trie[K]{
							'o': {
								end: true,
								childs: map[K]*trie[K]{
									'o': {
										end:    true,
										childs: map[K]*trie[K]{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := assert.New(t)
			tt.tr.put(tt.word)
			is.Equal(tt.wantTr, tt.tr)
		})
	}
}
