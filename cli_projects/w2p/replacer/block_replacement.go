package replacer

import (
	"bufio"
	"io"
)

// ReplaceSingleBlockBytes will return a []byte with the updated content of the reader if needs to make a replacement
// in it, otherwise will not return the bytes, if an error occurs during the reading it will stop and return an the error
func ReplaceSingleBlockBytes(reader io.Reader, replacer Replacer) ([]byte, bool, error) {
	data := make([]byte, 0)
	scanner := bufio.NewScanner(reader)

	replaced := &onlyTrue{}
	for scanner.Scan() {
		line := scanner.Bytes()

		if replaced.isTrue {
			data = addLine(data, line)
			continue
		}

		content, rep, err := replacer.replace(scanner, line)
		if err != nil {
			return nil, false, err
		}

		data = addLine(data, content)
		replaced.set(rep)
	}

	if err := scanner.Err(); err != nil {
		return nil, false, err
	}

	if replaced.isTrue {
		// remove the last '\n' character
		return data[:len(data)-1], replaced.isTrue, nil
	}

	// do not return the bytes if there's not a write to do
	return nil, false, nil
}
