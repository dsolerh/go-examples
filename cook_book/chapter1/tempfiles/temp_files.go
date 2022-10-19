package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WorkWithTemp will give some basic patterns for working
// with temporary files and directories
func WorkWithTemp() error {
	// If you need for a temporary place to store files with
	// the same name ie. template1-10.html a temp directory
	// is a good way to approach it, the first argument
	// being blank means it will create the directory
	// in the location returned by os.TempDir()
	fmt.Printf("os.TempDir(): %v\n", os.TempDir())
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	// this will delete everything inside the temp file
	// when this function returns if you want to do this
	// latter, be sure to return the directory name to the
	// calling function
	defer os.RemoveAll(t)

	// the directory must exist to create the tempfile
	// t is an *os.File object.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())
	// normally we'd delete the temporary file here, but
	// because we're placing it in a temp directory, it
	// gets cleaned up by the earlier defer

	return nil
}
