package main

import (
	"aoc/common/collections"
	"bufio"
	"bytes"
)

func solutionPart1(scanner *bufio.Scanner) int64 {
	var total int64

	pattern := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.Equal(line, []byte("")) {
			// calculate the reflections
			total += reflection(pattern)
			// clean the pattern
			pattern = nil
			continue
		}
		pattern = append(pattern, line)
	}

	return total
}

func reflection(pattern [][]byte) int64 {
	if v := verticalReflection(pattern); v != 0 {
		return v
	}
	if h := horizontalReflection(pattern); h != 0 {
		return h * 100
	}
	return 0
}

func verticalReflection(pattern [][]byte) int64 {
	for i := 0; i < len(pattern[0])-1; i++ {
		if hasSymmetry(pattern[0], i) {
			if collections.Every(pattern[1:], func(line []byte) bool { return hasSymmetry(line, i) }) {
				return int64(i + 1)
			}
		}
	}
	return 0
}

func horizontalReflection(pattern [][]byte) int64 {
	grid := NewGridIter(pattern)
	col0 := grid.Col(0)
	colRange := collections.Range(1, len(pattern[0]))
	for i := 0; i < len(pattern)-1; i++ {
		if hasSymmetry(col0, i) {
			if collections.Every(colRange, func(col int) bool { return hasSymmetry(grid.Col(col), int(i)) }) {
				return int64(i + 1)
			}
		}
	}
	return 0
}

func hasSymmetry(patternLine []byte, pos int) bool {
	for i := 0; i < len(patternLine); i++ {
		nextPos := pos + 1 + i
		prevPos := pos - i
		if prevPos < 0 || nextPos >= len(patternLine) {
			return true
		}
		if patternLine[prevPos] != patternLine[nextPos] {
			return false
		}
	}
	return true
}

// #.##..##.
// ^       ^
// #.##..##.
//  ^      ^
// #.##..##.
//   ^    ^
// #.##..##.
//   ^    ^

// #.##..##.
// ^       ^
// #.##..##.
// ^      ^ R
// #.##..##.
//  ^    ^  B
// #.##..##.
//   ^   ^  L
// #.##..##.
//    ^ ^	R
// #.##..##.
//     ^^	L

// ##......#
// ^       ^
// ##......#
//  ^     ^ B
// ##......#
//   ^    ^ L
// ##......#
//    ^  ^  B
// ##......#
//     ^^   B

// ##......#
// ^       ^
// ##......#
//  ^     ^ B
// ##......#
//  ^    ^  R
// ##......#
//   ^   ^  B
// ##......#
//    ^ ^   B
// ##......#
//     ^^   B

// ..#.##..##.
// ^.........^
// ..#.##..##.
// .^.......^. B
// ..#.##..##.
// ..^......^. L
// ..#.##..##.
// ...^....^.. B
// ..#.##..##.
// ...^...^... R
// ..#.##..##.
// ....^.^.... B
// ..#.##..##.
// .....^^.... L

// ..#.##..##.
// ....^^.....
// ..#.##..##.
// ...^..^....
// ..#.##..##.
// ..^....^...
// ..#.##..##.
// ......^^...
