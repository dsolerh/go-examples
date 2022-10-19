package bytestring

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some tricks for initializing bytes
// Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	// we'll start with a string encoded into raw bytes
	rawBytes := []byte(rawString)
	// there are a number of ways to create a buffer from
	// the raw bytes or from the original string
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	// alternatively
	b = bytes.NewBuffer(rawBytes)
	// and directly from the string
	b = bytes.NewBufferString(rawString)
	return b
}

// ToString is an example of consuming the io.Reader completley and returning a string
func ToString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
