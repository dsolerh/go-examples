package parser

import "errors"

var ErrEndOfData = errors.New("data terminated early")

type Output struct {
	Error     error
	Out       any
	Remaining []byte
}

type ParseFn = func(data []byte) Output

type ParsedData struct {
	err       error
	parsed    []any
	remaining []byte
}

func ParseInput(data []byte, parsers ...ParseFn) *ParsedData {
	parsedAll := make([]any, 0, len(parsers))
	for _, parser := range parsers {
		if len(data) == 0 {
			return &ParsedData{err: ErrEndOfData}
		}
		out := parser(data)
		if out.Error != nil {
			return &ParsedData{err: out.Error}
		}

		if out.Out != nil {
			parsedAll = append(parsedAll, out.Out)
		}
		data = out.Remaining
	}
	return &ParsedData{
		parsed:    parsedAll,
		remaining: data,
	}
}
