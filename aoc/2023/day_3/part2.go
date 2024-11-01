package main

import (
	"aoc/common/searcher"
	"bufio"
	"strconv"
)

type position struct {
	row int
	col int
}

type number struct {
	value int64
	start int
	end   int
}

func solutionPart2(scanner *bufio.Scanner) int64 {
	// per line numbers
	numbers := make([][]number, 0)
	gears := make([]position, 0)
	rowCounter := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		lineNumbers := make([]number, 0)
		s := searcher.NewLineSearcher(line, searcher.Numbers, searchGears)

		for s.Next() {
			data, at := s.Item()
			if isGear(data) {
				gears = append(gears, position{rowCounter, at})
			} else {
				lineNumbers = append(lineNumbers, number{value: parseNumber(data), start: at, end: at + len(data) - 1})
			}
		}

		rowCounter++
		numbers = append(numbers, lineNumbers)
	}

	var total int64
	for _, gear := range gears {
		var partsFound []int64

		// look for in the current row
		// this can produce at most 2 numbers, so cannot check for early
		// exit here
		partsFound = append(partsFound, findAdjacentPartNumbers(numbers[gear.row], gear.col)...)

		// check if there's a previous row
		if prevRow := gear.row - 1; prevRow >= 0 {
			partsFound = append(partsFound, findAdjacentPartNumbers(numbers[prevRow], gear.col)...)
			// here there can be an early check for the parts found
			if len(partsFound) > 2 {
				// stop looking for more parts as this gear is already not valid
				continue
			}
		}

		// check if there's a next row
		if nextRow := gear.row + 1; nextRow < len(numbers) {
			partsFound = append(partsFound, findAdjacentPartNumbers(numbers[nextRow], gear.col)...)
		}

		// if there's exactly two part numbers then add their product
		if len(partsFound) == 2 {
			total += (partsFound[0] * partsFound[1])
		}
	}

	return total
}

func findAdjacentPartNumbers(numbers []number, col int) []int64 {
	found := make([]int64, 0, 2)
	for _, numb := range numbers {
		if (col >= numb.start-1 && col <= numb.start+1) || (col >= numb.end-1 && col <= numb.end+1) {
			found = append(found, numb.value)
		}
	}
	return found
}

func parseNumber(data []byte) int64 {
	val, _ := strconv.Atoi(string(data))
	return int64(val)
}

func isGear(data []byte) bool {
	return data[0] == '*'
}

func searchGears(data []byte) []byte {
	if isGear(data) {
		return []byte{'*'}
	}
	return nil
}
