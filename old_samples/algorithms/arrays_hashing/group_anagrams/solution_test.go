package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example 1",
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{
				{"hat"},
				{"act", "cat"},
				{"pots", "tops", "stop"},
			},
		},
		{
			name: "example 2",
			strs: []string{"x"},
			want: [][]string{{"x"}},
		},
		{
			name: "example 3",
			strs: []string{""},
			want: [][]string{{""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.strs, hashv2)
			is := assert.New(t)
			is.ElementsMatch(got, tt.want)
		})
	}
}
