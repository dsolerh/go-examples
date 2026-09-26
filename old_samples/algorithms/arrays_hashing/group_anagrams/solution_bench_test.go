package groupanagrams

import (
	"testing"
)

var anaGroup [][]string

func BenchmarkGroupAnagrams(b *testing.B) {
	b.Run("with hashv1", func(b *testing.B) {
		var ag [][]string
		for b.Loop() {
			ag = GroupAnagrams(
				[]string{
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
				},
				hashv1,
			)
		}
		anaGroup = ag
	})
	b.Run("with hashv2", func(b *testing.B) {
		var ag [][]string
		for b.Loop() {
			anaGroup = GroupAnagrams(
				[]string{
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"act", "pots", "tops", "cat", "stop", "hat",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
					"safadsfsadfasdfasdfsadfsvcvdvdfgsdfsdfsdda",
				},
				hashv2,
			)
		}
		anaGroup = ag
	})
}
