package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s, sep := "a:b:c", ":"
	words := strings.Split(s, sep)
	if got, want := len(words), 3; got != want {
		t.Errorf("Split(%q,%q) returned %d words, want %d", s, sep, got, want)
	}
}

func TestTSplit(t *testing.T) {
	testCases := []struct {
		s    string
		sep  string
		want int
	}{
		{
			s:    "a:b:c",
			sep:  ":",
			want: 3,
		},
		{
			s:    "a,b:c",
			sep:  ":",
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("Spliting: [ %q ] with %q", tC.s, tC.sep), func(t *testing.T) {
			words := strings.Split(tC.s, tC.sep)
			if got := len(words); got != tC.want {
				t.Errorf("Split(%q, %q) returned %d words, want %d", tC.s, tC.sep, got, tC.want)
			}
		})
	}
}
