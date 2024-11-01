package replacer

import (
	"bufio"
	"bytes"
	"io"
)

type Replacements map[string][]byte

func (r Replacements) replace(line []byte) ([]byte, bool) {
	for oldS, newB := range r {
		oldB := []byte(oldS)
		if bytes.Contains(line, oldB) {
			return bytes.Replace(line, oldB, newB, 1), true
		}
	}
	return line, false
}

func LoadReplacements(reader io.Reader) (Replacements, error) {
	replacements := make(Replacements, 0)
	scanner := bufio.NewScanner(reader)
external:
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, useDeclarationStart) {
			for scanner.Scan() {
				line = scanner.Bytes()

				// check if for the end of the use block
				if bytes.Equal(line, multiImportPrefixEnd) {
					break external
				}

				// parse the line data
				lineSplited := bytes.SplitN(line, replaceDeclaration, 2)
				if len(lineSplited) < 2 {
					continue
				}

				newPath := bytes.TrimSpace(lineSplited[1])
				if len(newPath) == 0 {
					continue
				}

				oldPath, _ := bytes.CutPrefix(bytes.TrimSpace(lineSplited[0]), currentFolderSegment)
				replacements[string(oldPath)] = newPath
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return replacements, nil
}
