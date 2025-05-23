package main

import (
	"aoc/data"
	"bufio"
	"bytes"
	"fmt"
	"slices"
)

func main() {
	data, err := data.GetData(2023, 1)
	if err != nil {
		fmt.Printf("getting data err: %v\n", err)
		return
	}
	total := calculateSum(bufio.NewScanner(data), computeNumberV2)
	fmt.Printf("total: %v\n", total)
}

func calculateSum(scanner *bufio.Scanner, computeNumberFn func([]byte) int64) int64 {
	var total int64

	for scanner.Scan() {
		total += computeNumberFn(scanner.Bytes())
	}
	return total
}

func computeNumberV1(line []byte) int64 {
	numbers := make([]byte, 0)

	for current := range line {
		if isDigit(line[current]) {
			numbers = append(numbers, line[current])
		}
	}

	if len(numbers) == 0 {
		return 0
	}

	first := toDigit(numbers[0])
	last := toDigit(numbers[len(numbers)-1])
	return 10*first + last
}

var NumWords = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

func computeNumberV2(line []byte) int64 {
	first := firstNumber(line)
	last := lastNumber(line)
	return 10*first + last
}

func firstNumber(line []byte) int64 {
	for current := range line {
		if current != 0 {
			currentWord := line[:current]
			index := slices.IndexFunc(NumWords, func(numWord []byte) bool {
				return bytes.HasSuffix(currentWord, numWord)
			})
			if index != -1 {
				return int64(index + 1) // add 1 to the index since the NumWords are ordered and start at 0 index
			}
		}
		if isDigit(line[current]) {
			return toDigit(line[current])
		}
	}
	return 0
}

func lastNumber(line []byte) int64 {
	for current := len(line) - 1; current >= 0; current-- {
		if current != len(line)-1 {
			currentWord := line[current:]
			index := slices.IndexFunc(NumWords, func(numWord []byte) bool {
				return bytes.HasPrefix(currentWord, numWord)
			})
			if index != -1 {
				return int64(index + 1) // add 1 to the index since the NumWords are ordered and start at 0 index
			}
		}
		if isDigit(line[current]) {
			return toDigit(line[current])
		}
	}
	return 0
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func toDigit(b byte) int64 {
	return int64(b - '0')
}
