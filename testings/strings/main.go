package main

// Group Anagrams
// You are given an array of strings. Your task is to group the anagrams together. An anagram is defined as two words that can be formed by rearranging the letters of one another.
// Write a function groupAnagrams(strs) that takes a list of strings and returns a list of lists, where each sublist contains the anagrams grouped together.
// Example:
// Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "teacup", "peacut"}
	anagramsGrouped := map[string][]string{}
	for _, word := range strs {
		anaKey := getAnaKey(word)
		anagramsGrouped[anaKey] = append(anagramsGrouped[anaKey], word)
	}

	anagrams := [][]string{}
	for _, groupedAnagrams := range anagramsGrouped {
		anagrams = append(anagrams, groupedAnagrams)
	}

	fmt.Printf("anagrams: %v\n", anagrams)
}

func getAnaKey(word string) string {
	// sorted
	// "tea cup" === "peacut"

	letters := []int{}
	for _, letter := range word {
		// if letter != ' ' {
		//
		// }
		letters = append(letters, int(letter))
	}
	sort.Ints(letters)

	str := strings.Builder{}
	for _, letter := range letters {
		str.WriteRune(rune(letter))
	}
	return str.String()
}

// // names := []string{"daniel", "  daniel  ", "%%%%daniel%%%", " % %%% daniel"}
// 	// for _, name := range names {
// 	// 	trimName := string(bytes.TrimLeft([]byte(name), "% "))
// 	// 	fmt.Printf("name: [%v]\n", name)
// 	// 	fmt.Printf("trimName: [%v]\n", trimName)
// 	// 	fmt.Println()
// 	// }
// 	a := []byte{91, 34, 34, 44, 49, 93}
// 	fmt.Printf("a: %v\n", a)
// 	fmt.Printf("a: %s\n", a)
// 	time.Now()
// §

//
// Imagine you are developing a multithreaded application for a banking system that allows customers to access and manage their bank accounts online. In this application, multiple users can simultaneously attempt to withdraw or deposit money into the same account.
// How would you ensure data integrity?
// What challenges might arise when multiple users try to access and modify the same bank account concurrently?
// Can you provide an example of a situation in which some users might find it difficult to complete their transactions in a timely manner?

// account current-balance
// 1		12		(+20)(-30)
// 1		2

// 1		12
// 1		+20
// 1		-30
// 1		-100
// -----------------
// 1		2 (emit an event)
// -----------------
// 1		-98
