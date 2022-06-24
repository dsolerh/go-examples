package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkdir(cwd, name string) (string, error) {
	migDir := filepath.Join(cwd, name)
	info, err := os.Stat(migDir)
	if err != nil && os.IsNotExist(err) {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("%s is not a directory", migDir)
	}
	return migDir, err
}
