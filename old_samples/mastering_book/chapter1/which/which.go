package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("An argument is required!")
		return
	}
	for _, file := range arguments[1:] {
		path := os.Getenv("PATH")
		pathSplit := filepath.SplitList(path)
		for _, dir := range pathSplit {
			fullExecPath, found := getExec(dir, file)
			if found {
				fmt.Println(fullExecPath)
			}
		}
	}
}

func getExec(dir, file string) (string, bool) {
	fullPath := filepath.Join(dir, file)
	// check if exist
	fileInfo, err := os.Stat(fullPath)
	if err == nil {
		mode := fileInfo.Mode()
		if mode.IsRegular() {
			if mode&0111 != 0 {
				return fullPath, true
			}
		}
	}
	return "", false
}
