package main

import (
	"aoc/common/collections"
	"aoc/common/parser"
	"bufio"
	"slices"
)

func solutionPart1(scanner *bufio.Scanner) int64 {
	var total int64

	for scanner.Scan() {
		line := scanner.Bytes()
		parsed := parser.ParseInput(line,
			parser.UntilChar(' '),
			parser.Multiples(parser.UntilEnd(), parser.Number, parser.Discard(parser.Char(','))),
		)

		springs, _ := parsed.Bytes(0)
		damagedSprings, _ := parsed.Numbers(1)

		total += recursive(springs, damagedSprings)
	}

	return total
}

func recursive(springs []byte, damagedSprings []int64) int64 {
	if len(damagedSprings) == 0 {
		// cannot have a spring broken remaining
		if collections.Some(springs, isBroken) {
			return 0
		}
		return 1 // found a convination
	}

	if len(springs) == 0 {
		return 0
	}

	switch springs[0] {
	case '?':
		// check the paths
		spring := damagedSprings[0]
		section := collections.Take(springs, int(spring))
		if len(section) == 0 {
			// if cannot take the whole section then return
			return 0
		}

		if !collections.Every(section, isUnknownOrBroken) {
			// this is not a broken section, move along
			amount := recursive(springs[1:], damagedSprings)
			return amount
		}

		// check if there's a position after the last
		if int(spring) == len(springs) {
			amount := recursive(nil, damagedSprings[1:]) // recur
			return amount
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			amount := recursive(springs[1:], damagedSprings) // recur (but dont remove the damaged)
			return amount
		}

		// move till the end of the section+1
		amount := recursive(springs[spring+1:], damagedSprings[1:]) + recursive(springs[1:], damagedSprings)
		return amount

	case '#':
		// it's a broken gear continue from it's end
		spring := damagedSprings[0]
		section := collections.Take(springs, int(spring))
		if len(section) == 0 {
			// if cannot take the whole section then return
			return 0
		}

		if !collections.Every(section, isUnknownOrBroken) {
			// this is not a broken section, stop searching
			return 0
		}

		// check if there's a position after the last
		if int(spring) == len(springs) {
			amount := recursive(nil, damagedSprings[1:]) // recur
			return amount
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			return 0 // stop
		}

		// move till the end of the section+1
		amount := recursive(springs[spring+1:], damagedSprings[1:])
		return amount

	default:
		idx := slices.IndexFunc(springs, isUnknownOrBroken)
		if idx == -1 {
			return recursive(nil, damagedSprings)
		}
		return recursive(springs[idx:], damagedSprings)
	}
}

func isBroken(b byte) bool {
	return b == '#'
}

func isUnknownOrBroken(b byte) bool {
	return b == '?' || b == '#'
}
