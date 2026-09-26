package main

import (
	"fmt"
	"os"
	"time"
)

const dateFormat = "02 January 2006 15:04 MST"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	d, err := time.Parse(dateFormat, dateString)
	if err == nil {
		loc, _ := time.LoadLocation("Local")
		fmt.Printf("Local Time:\t\t%s\n", d.In(loc))

		loc, _ = time.LoadLocation("UTC")
		fmt.Printf("UTC Time:\t\t%s\n", d.In(loc))

		loc, _ = time.LoadLocation("America/New_York")
		fmt.Printf("New York Time:\t\t%s\n", d.In(loc))
	} else {
		fmt.Printf("Invalid date format\nTry: \"%s\"\n", dateFormat)
	}
}
