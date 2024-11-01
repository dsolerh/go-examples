package parser

import "fmt"

func Multiples(parser ParseFn, innerParsers ...ParseFn) ParseFn {
	return func(data []byte) Output {
		outOut := parser(data)
		if outOut.Error != nil {
			return Output{Error: fmt.Errorf("error while parsing multiples: %w", outOut.Error)}
		}

		innerData, ok := outOut.Out.([]byte)
		if !ok {
			return Output{Error: fmt.Errorf("invalid intermediate data format: %T", outOut.Out)}
		}

		innerParserIndex := 0
		innerParsersCount := len(innerParsers)
		parsed := make([]any, 0)

		for len(innerData) > 0 {
			out := innerParsers[innerParserIndex%innerParsersCount](innerData)
			if out.Error != nil {
				return Output{Error: fmt.Errorf("error while parsing inner parser: %w", out.Error)}
			}

			if out.Out != nil {
				parsed = append(parsed, out.Out)
			}
			innerData = out.Remaining
			innerParserIndex++
		}

		return Output{
			Out:       parsed,
			Remaining: outOut.Remaining,
		}
	}
}
