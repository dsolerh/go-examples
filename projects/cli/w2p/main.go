package main

import (
	"fmt"
)

func main() {

}

func checkErr(msg string, err error, args ...any) {
	if err != nil {
		panic(fmt.Sprintf(msg, append([]any{err}, args...)))
	}
}

// func replaceFileImports(fname string, replacements replacements) {
// 	f, err := os.Open(fname)
// 	checkErr("error opening the file: %v", err)
// 	defer f.Close()

// 	data, replace, err := replaceImports(f, replacements)

// 	checkErr("error reading the file: %v", err)

// 	if replace {
// 		err = os.WriteFile(fname, data, 0o600)
// 		checkErr("could not write the file: %v", err)
// 	}
// }

// func replaceFileModuleName(fname string, replacements replacements) {
// 	f, err := os.Open(fname)
// 	checkErr("error opening the file: %v", err)
// 	defer f.Close()

// 	data, replace, err := replaceImports(f, replacements)

// 	checkErr("error reading the file: %v", err)

// 	if replace {
// 		err = os.WriteFile(fname, data, 0o600)
// 		checkErr("could not write the file: %v", err)
// 	}
// }
