package main

import "bufio"

func solutionPart2(scanner *bufio.Scanner) int64 {
	var total int64

	for scanner.Scan() {
		info := parseLine(scanner.Bytes())

		var minimalRed, minimalGreen, minimalBlue int
		for _, extraction := range info.extractions {
			if extraction.red > minimalRed {
				minimalRed = extraction.red
			}
			if extraction.green > minimalGreen {
				minimalGreen = extraction.green
			}
			if extraction.blue > minimalBlue {
				minimalBlue = extraction.blue
			}
		}
		total += (int64(minimalRed) * int64(minimalGreen) * int64(minimalBlue))
	}

	return total
}
