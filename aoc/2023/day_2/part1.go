package main

import (
	"bufio"
	"slices"
)

func solutionPart1(scanner *bufio.Scanner) int64 {
	var total int64

	for scanner.Scan() {
		info := parseLine(scanner.Bytes())

		index := slices.IndexFunc(info.extractions, func(e extraction) bool {
			return e.red > MaxRed || e.green > MaxGreen || e.blue > MaxBlue
		})

		if index == -1 {
			total += info.id
		}
	}

	return total
}
