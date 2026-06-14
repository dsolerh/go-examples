package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments")
	}
	file1 := os.Args[1]
	file2 := os.Args[2]

	data, err := os.ReadFile(file1)
	if err != nil {
		log.Fatal(err)
	}

	modFile1, err := modfile.Parse(file1, data, nil)
	if err != nil {
		log.Fatal(err)
	}

	data, err = os.ReadFile(file2)
	if err != nil {
		log.Fatal(err)
	}

	modFile2, err := modfile.Parse(file2, data, nil)
	if err != nil {
		log.Fatal(err)
	}

	// // Access module information
	// fmt.Println("Module:", modFile.Module.Mod.Path)
	// fmt.Println("Go version:", modFile.Go.Version)

	// // List requirements
	// for _, req := range modFile.Require {
	// 	fmt.Printf("Require: %s %s\n", req.Mod.Path, req.Mod.Version)
	// }
}
