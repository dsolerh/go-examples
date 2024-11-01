package replacer

import (
	"bufio"
)

type onlyTrue struct {
	isTrue bool
}

func (o *onlyTrue) set(b bool) {
	if !o.isTrue {
		o.isTrue = b
	}
}

type Replacer interface {
	replace(s *bufio.Scanner, currentLine []byte) ([]byte, bool, error)
}
