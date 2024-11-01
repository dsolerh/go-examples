package searcher

import "aoc/common/digits"

func Numbers(data []byte) []byte {
	if !digits.IsDigit(data[0]) {
		return nil
	}

	index := 1
	for ; index < len(data); index++ {
		if !digits.IsDigit(data[index]) {
			break
		}
	}

	return data[:index]
}
