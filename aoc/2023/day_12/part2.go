package main

import (
	"aoc/common/collections"
	"aoc/common/ds"
	"aoc/common/parser"
	"bufio"
	"fmt"
	"slices"
)

func solutionPart2(scanner *bufio.Scanner) int64 {
	var total int64

	for scanner.Scan() {
		line := scanner.Bytes()
		parsed := parser.ParseInput(line,
			parser.UntilChar(' '),
			parser.Multiples(parser.UntilEnd(), parser.Number, parser.Discard(parser.Char(','))),
		)

		springs, _ := parsed.Bytes(0)
		damagedSprings, _ := parsed.Numbers(1)

		amount := recursiveWithCache(
			collections.Join(collections.Repeat(springs, 5), '?'),
			collections.Flat(collections.Repeat(damagedSprings, 5)),
			map[string]int64{},
		)
		total += amount
	}

	return total
}

type args struct {
	springs        []byte
	damagedSprings []int64
}

func iterative(springs []byte, damagedSprings []int64) int64 {
	var total int64

	stack := ds.NewStack[args]()
	stack.Push(args{springs, damagedSprings})

	for !stack.Empty() {
		total += procesArgs(stack)
	}

	return total
}

func procesArgs(stack *ds.Stack[args]) int64 {
	a := stack.Pop()
	springs := a.springs
	damagedSprings := a.damagedSprings

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
			stack.Push(args{springs[1:], damagedSprings})
			return 0
		}

		// check if there's a position after the last
		if int(spring) == len(springs) {
			stack.Push(args{nil, damagedSprings[1:]})
			return 0
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			stack.Push(args{springs[1:], damagedSprings})
			return 0
		}

		// move till the end of the section+1
		stack.Push(args{springs[spring+1:], damagedSprings[1:]})
		stack.Push(args{springs[1:], damagedSprings})
		return 0

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
			stack.Push(args{nil, damagedSprings[1:]})
			return 0
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			return 0
		}

		// move till the end of the section+1
		stack.Push(args{springs[spring+1:], damagedSprings[1:]})
		return 0

	default:
		stack.Push(args{springs[1:], damagedSprings})
		return 0
	}
}

func recursiveWithCache(springs []byte, damagedSprings []int64, cache map[string]int64) int64 {
	key := cacheKey(springs, damagedSprings)
	if amount, exist := cache[key]; exist {
		return amount
	}

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
			amount := recursiveWithCache(springs[1:], damagedSprings, cache)
			cache[key] = amount
			return amount
		}

		// check if there's a position after the last
		if int(spring) == len(springs) {
			amount := recursiveWithCache(nil, damagedSprings[1:], cache) // recur
			cache[key] = amount
			return amount
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			amount := recursiveWithCache(springs[1:], damagedSprings, cache) // recur (but dont remove the damaged)
			cache[key] = amount
			return amount
		}

		// move till the end of the section+1
		amount := recursiveWithCache(springs[spring+1:], damagedSprings[1:], cache) + recursiveWithCache(springs[1:], damagedSprings, cache)
		cache[key] = amount
		return amount

	case '#':
		// it's a broken gear continue from it's end
		spring := damagedSprings[0]
		section := collections.Take(springs, int(spring))
		if len(section) == 0 {
			// if cannot take the whole section then return
			cache[key] = 0
			return 0
		}

		if !collections.Every(section, isUnknownOrBroken) {
			// this is not a broken section, stop searching
			cache[key] = 0
			return 0
		}

		// check if there's a position after the last
		if int(spring) == len(springs) {
			amount := recursiveWithCache(nil, damagedSprings[1:], cache) // recur
			cache[key] = amount
			return amount
		}

		// check if the position after the last holds a broken
		if isBroken(springs[spring]) {
			cache[key] = 0
			return 0 // stop
		}

		// move till the end of the section+1
		amount := recursiveWithCache(springs[spring+1:], damagedSprings[1:], cache)
		cache[key] = amount
		return amount

	default:
		idx := slices.IndexFunc(springs, isUnknownOrBroken)
		if idx == -1 {
			return recursiveWithCache(nil, damagedSprings, cache)
		}
		return recursiveWithCache(springs[idx:], damagedSprings, cache)
	}
}

func cacheKey(springs []byte, damagedSprings []int64) string {
	return fmt.Sprintf("%s%v", springs, damagedSprings)
}
