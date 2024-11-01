package main

import (
	"aoc/common/ds"
	"aoc/common/parser"
	"bufio"
)

func solutionPart1(scanner *bufio.Scanner) int64 {
	var total int64

	parsers := []parser.ParseFn{
		parser.Discard(parser.UntilChar(':')),
		parser.Multiples(parser.UntilChar('|'), parser.Discard(parser.Space), parser.Number),
		parser.Multiples(parser.UntilEnd(), parser.Discard(parser.Space), parser.Number),
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		parsed := parser.ParseInput(line, parsers...)

		winningNumbers, _ := parsed.Numbers(0)
		winningNumbersSet := ds.NewSet(winningNumbers)
		ownedNumbers, _ := parsed.Numbers(1)

		var cardWorth int64
		for _, owned := range ownedNumbers {
			if winningNumbersSet.Has(owned) {
				if cardWorth == 0 {
					cardWorth = 1
				} else {
					cardWorth *= 2
				}
			}
		}

		total += cardWorth
	}

	return total
}
