package main

import (
	"aoc/common/digits"
	"bufio"
	"strconv"
)

func solutionPart1(scanner *bufio.Scanner) int64 {
	// create the schematics matrix
	schematic := make([][]byte, 0)
	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())
		schematic = append(schematic, line)
	}

	var total int64
	for row := 0; row < len(schematic); row++ {
		currentLine := schematic[row]
		start := 0
		for start < len(currentLine) {
			// if the current byte is a digit it marks the beginning of a number
			if digits.IsDigit(currentLine[start]) {
				current := start + 1

				// iterate till the current value is no longer a digit
				for ; current < len(currentLine); current++ {
					if !digits.IsDigit(currentLine[current]) {
						break
					}
				}
				// the value of current here is the +1 the value at which the previous condition
				// was met
				current--

				var value int64
				if current == start {
					// check all possible adjacent directions from start
					if adjacentSymbol(schematic, row, start) {
						value = digits.ToDigit(currentLine[start])
					}
				} else {
					// check all possible adjacent directions from start and current
					if adjacentSymbol(schematic, row, start) || adjacentSymbol(schematic, row, current) {
						i, _ := strconv.Atoi(string(currentLine[start : current+1]))
						value = int64(i)
					}
				}

				total += value

				// update start
				start = current + 1
			} else {
				// increase the start position to keep searching
				start++
			}
		}
	}

	return total
}

func adjacentSymbol(schematic [][]byte, row, col int) bool {
	if prevRow := row - 1; prevRow >= 0 {
		if digits.IsSymbol(schematic[prevRow][col]) {
			return true
		}
		if prevCol := col - 1; prevCol >= 0 && digits.IsSymbol(schematic[prevRow][prevCol]) {
			return true
		}
		if nextCol := col + 1; nextCol < len(schematic[prevRow]) && digits.IsSymbol(schematic[prevRow][nextCol]) {
			return true
		}
	}

	if nextRow := row + 1; nextRow < len(schematic) {
		if digits.IsSymbol(schematic[nextRow][col]) {
			return true
		}
		if prevCol := col - 1; prevCol >= 0 && digits.IsSymbol(schematic[nextRow][prevCol]) {
			return true
		}
		if nextCol := col + 1; nextCol < len(schematic[nextRow]) && digits.IsSymbol(schematic[nextRow][nextCol]) {
			return true
		}
	}

	if prevCol := col - 1; prevCol >= 0 && digits.IsSymbol(schematic[row][prevCol]) {
		return true
	}
	if nextCol := col + 1; nextCol < len(schematic[row]) && digits.IsSymbol(schematic[row][nextCol]) {
		return true
	}

	return false
}
