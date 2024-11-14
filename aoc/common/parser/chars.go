package parser

import "fmt"

const (
	NoChar = "invalid char at (%d) in data (%s)"
)

func Char(ch byte) ParseFn {
	return func(data []byte) Output {
		if data[0] != ch {
			return Output{Error: fmt.Errorf(NoBlankSpace, 0, data)}
		}

		index := 1
		for ; index < len(data); index++ {
			if data[index] != ' ' {
				break
			}
		}

		return Output{
			Out:       data[:index],
			Remaining: data[index:],
		}
	}
}
