package parser

import (
	"aoc/common/digits"
	"fmt"
)

type number = int64

const (
	NonNumericalDigitAt = "the value (%c) at (%d) in data (%s) is non numerical"
)

// Number takes digits from the beginig of the byte array till there's
// a non numerical character or the data is over, then returns the numerical value
// and the remaning bytes if any. If no numerical digit is found at the beguining of
// the data then an error is returned
func Number(data []byte) Output {
	if !digits.IsDigit(data[0]) {
		return Output{Error: fmt.Errorf(NonNumericalDigitAt, data[0], 0, data)}
	}

	index := 1
	for ; index < len(data); index++ {
		if !digits.IsDigit(data[index]) {
			break
		}
	}

	return Output{
		Out:       ParseNumber(data[:index]),
		Remaining: data[index:],
	}
}

func ParseNumber(data []byte) number {
	if len(data) >= 19 {
		panic("cannot parse numbers bigger than 64 bits")
	}
	var num number
	for _, ch := range data {
		num = 10*num + digits.ToDigit(ch)
	}
	return num
}
