package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trie_Has(t *testing.T) {
	tests := []struct {
		name string
		tr   trie
		word string
		want bool
	}{
		{
			name: "the word is contained in the trie",
			tr: trie{
				'f': {
					'o': {
						'o': {
							'l':  {},
							Null: {},
						},
					},
				},
			},
			word: "foo",
			want: true,
		},
		{
			name: "the word is not contained in the trie (is there but not intencionaly)",
			tr: trie{
				'f': {
					'o': {
						'o': {
							'l':  {},
							Null: {},
						},
					},
				},
			},
			word: "fo",
			want: false,
		},
		{
			name: "the word is not contained in the trie",
			tr: trie{
				'f': {
					'o': {
						'o': {
							'l':  {},
							Null: {},
						},
					},
				},
			},
			word: "footer",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Has(tt.word); got != tt.want {
				t.Errorf("trie.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trie_Put(t *testing.T) {
	tests := []struct {
		tr     trie
		wantTr trie
		name   string
		word   string
	}{
		{
			name:   "should add the string to an empty trie",
			word:   "foo",
			tr:     trie{},
			wantTr: trie{'f': {'o': {'o': {Null: nil}}}},
		},
		{
			name:   "should add the string to an existing trie",
			word:   "foz",
			tr:     trie{'f': {'o': {'o': {Null: nil}}}},
			wantTr: trie{'f': {'o': {'o': {Null: nil}, 'z': {Null: nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := assert.New(t)
			tt.tr.Put(tt.word)
			is.Equal(tt.wantTr, tt.tr)
		})
	}
}
