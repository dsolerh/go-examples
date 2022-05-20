package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const CSVFILE = "data.csv"

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list <arguments>")
		return
	}

	// If the the file does not exist, create one
	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}
	createIndex()

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		record, err := parseRecord([]string{arguments[2], arguments[3], t, time.Now().Format(ISOlayout)})
		if err != nil {
			fmt.Println(err)
			return
		}
		insertRecord(record)
	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		err := deleteRecord(t)
		if err != nil {
			fmt.Println(err)
		}
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Number")
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		found := searchRecord(t)
		if found == nil {
			fmt.Println("Number not found:", t)
			return
		}
		fmt.Println(*found)
	case "list":
		listRecord()
	default:
		fmt.Println("Not a valid option")
	}
	saveCSVFile(CSVFILE)
}
