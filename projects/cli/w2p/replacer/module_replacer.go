package replacer

import (
	"bufio"
	"bytes"
)

var _ Replacer = (*importReplacer)(nil)

type moduleReplacer struct {
	replacements Replacements
}

func NewModuleReplacer(r Replacements) Replacer {
	return &moduleReplacer{replacements: r}
}

// replace implements Replacer.
func (m *moduleReplacer) replace(s *bufio.Scanner, currentLine []byte) ([]byte, bool, error) {
	data := make([]byte, 0)

	replaced := false
	for s.Scan() {
		line := s.Bytes()

		if replaced {
			data = addLine(data, line)
			continue
		}

		if bytes.HasPrefix(line, moduleDeclaration) {
			// replace single import
			var rep bool
			line, rep = m.replacements.replace(line)
			data = addLine(data, line)
			if rep {
				replaced = rep
			}
			continue
		}

		data = addLine(data, line)
	}

	if err := s.Err(); err != nil {
		return nil, false, err
	}

	if replaced {
		return data, replaced, nil
	}

	// do not return the bytes if there's not a write to do
	return nil, false, nil
}
