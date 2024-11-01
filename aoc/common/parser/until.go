package parser

import "fmt"

const (
	CharNotFound = "the char (%c) was not found in data (%s)"
)

func UntilChar(char byte) ParseFn {
	return func(data []byte) Output {
		found := false
		index := 0
		for ; index < len(data); index++ {
			if data[index] == char {
				found = true
				break
			}
		}

		if !found {
			return Output{Error: fmt.Errorf(CharNotFound, char, data)}
		}

		return Output{
			Out:       data[:index],
			Remaining: data[index+1:],
		}
	}
}

func UntilEnd() ParseFn {
	return func(data []byte) Output {
		return Output{Out: data}
	}
}
