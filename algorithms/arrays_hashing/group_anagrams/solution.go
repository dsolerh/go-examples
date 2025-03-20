package groupanagrams

import (
	"maps"
	"slices"
)

// hashv1 sorts the string to get a unique representation of it
// is kind of ineficient cause the sorting will take at most 'log n'
func hashv1(str string) string {
	b := []byte(str)
	slices.Sort(b)
	return string(b)
}

// hashv2 hashes the string using the frequency of each character and assumes that
// there's only 26 characters  (so only letters 'a'-'z')
func hashv2(str string) string {
	var alphabet [26]byte
	for i := 0; i < len(str); i++ {
		ch := str[i]
		alphabet[int(ch-'a')] += 1
	}
	return string(alphabet[:])
}

func GroupAnagrams(strs []string, hfn func(string) string) [][]string {
	anaMap := make(map[string][]string)
	for _, str := range strs {
		s := hfn(str)
		anaMap[s] = append(anaMap[s], str)
	}
	return slices.Collect(maps.Values(anaMap))
}
