package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy copies data from in to out first directly,
// then using a buffer. It also writes to stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	// we write to out and stdout
	w := io.MultiWriter(out, os.Stdout)

	// a standard copy, it can be dangerous if a lot of data is involved
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)

	// buffered write using 64 byte chunks
	buff := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buff); err != nil {
		return err
	}

	// print a new line
	fmt.Println()
	return nil
}
