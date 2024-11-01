package replacer

import (
	"bufio"
	"bytes"
)

var _ Replacer = (*importReplacer)(nil)

type importReplacer struct {
	replacements Replacements
}

func NewImportReplacer(r Replacements) Replacer {
	return &importReplacer{replacements: r}
}

// Replace implements replacer.
func (i *importReplacer) replace(s *bufio.Scanner, currentLine []byte) ([]byte, bool, error) {
	data := make([]byte, 0)

	if bytes.HasPrefix(currentLine, multiImportPrefixStart) {
		// add the `import (` line
		data = addLine(data, currentLine)

		// check for multiple imports
		replaced := &onlyTrue{} // starts being false and cannot be reset to false
		for s.Scan() {
			line := s.Bytes()

			// check if for the end of the import block
			if bytes.Equal(line, multiImportPrefixEnd) {
				// add end of import `)`
				data = addLine(data, line)
				break
			}

			// replace the line import
			var rep bool
			line, rep = i.replacements.replace(line)
			data = addLine(data, line)

			replaced.set(rep)
		}

		if err := s.Err(); err != nil {
			return nil, false, err
		}

		if len(data) > 1 {
			// remove the last '\n' character
			data = data[:len(data)-1]
		}
		return data, replaced.isTrue, nil
	}

	if bytes.HasPrefix(currentLine, singleImportPrefix) {
		// replace single import
		l, r := i.replacements.replace(currentLine)
		return l, r, nil
	}

	return currentLine, false, nil
}
