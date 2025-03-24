package encodedecode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Encode(strs []string) string {
	builder := new(strings.Builder)
	for _, str := range strs {
		l := len(str)
		builder.WriteString(fmt.Sprintf("%d|%s", l, str))
	}
	return builder.String()
}

func isAsciiDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func Decode(str string) ([]string, error) {
	if !isAsciiDigit(str[0]) {
		return nil, errors.New("invalid string sequence to decode, must start with a digit")
	}
	curr := 1
	decoded := make([]string, 0)
	for curr < len(str) {
		char := str[curr]
		if isAsciiDigit(char) {
			curr++
		} else if char == '|' && isAsciiDigit(str[curr-1]) {
			n, _ := strconv.Atoi(str[:curr])
			decoded = append(decoded, str[curr+1:curr+1+n])
			str = str[curr+1+n:] // move the string slice forward
			curr = 0             // reset curr to 0
		} else {
			return nil, errors.New("error while decoding")
		}
	}
	return decoded, nil
}
