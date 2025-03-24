package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a: %v\n", a)
	rand.Shuffle(len(a), Swap(a))
	fmt.Printf("a: %v\n", a)

	p := rand.Perm(3)
	fmt.Printf("p: %v\n", p)
}

func Swap[S ~[]E, E any](s S) func(int, int) {
	return func(i, j int) { s[i], s[j] = s[j], s[i] }
}
