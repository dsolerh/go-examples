package parser

import "fmt"

const (
	NoBlankSpace = "invalid blank space at (%d) in data (%s)"
)

func Space(data []byte) Output {
	if data[0] != ' ' {
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
