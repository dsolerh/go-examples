package parser

func Discard(parse ParseFn) ParseFn {
	return func(data []byte) Output {
		out := parse(data)
		if out.Error != nil {
			return out
		}

		// discards the data
		out.Out = nil
		return out
	}
}
