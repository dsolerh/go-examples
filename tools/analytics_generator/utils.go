package main

import (
	"fmt"
	"strings"
	"unicode"
)

func must[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
	return val
}

func formatPascalCase(w string) string {
	b := strings.Builder{}
	upper := true
	for _, letter := range w {
		if upper {
			upper = false
			b.WriteRune(unicode.ToUpper(letter))
		} else if unicode.IsSymbol(letter) || unicode.IsPunct(letter) {
			upper = true
		} else {
			b.WriteRune(letter)
		}
	}
	return b.String()
}

func formatCamelCase(w string) string {
	b := strings.Builder{}
	upper := false
	for _, letter := range w {
		if upper {
			upper = false
			b.WriteRune(unicode.ToUpper(letter))
		} else if letter == rune('_') {
			upper = true
		} else {
			b.WriteRune(letter)
		}
	}
	return b.String()
}
