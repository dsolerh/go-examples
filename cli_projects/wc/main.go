package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, lines bool) int {
	scanner := bufio.NewScanner(r)
	if lines {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	// create a variable to count the words
	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
