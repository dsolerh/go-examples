package main

import (
	"fmt"
	"unicode"
)

func main() {
	symbols := []rune{'_', '@', '#', '$', '%', '&', '*', '+', '=', '<', '>', '|', '~', '§', '¶', '†', '‡'}

	for _, r := range symbols {
		fmt.Printf("'%c': IsSymbol=%t, IsPunct=%t\n", r, unicode.IsSymbol(r), unicode.IsPunct(r))
	}
}
