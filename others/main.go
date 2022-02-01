package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// start := time.Now()
	// fmt.Printf("PopCount(108): %v\n", PopCount(999999999999))
	// fmt.Printf("time.Since(start).Seconds(): %v\n", time.Since(start).Seconds())

	// start = time.Now()
	// fmt.Printf("PopCountShift(108): %v\n", PopCountShift(999999999999))
	// fmt.Printf("time.Since(start).Seconds(): %v\n", time.Since(start).Seconds())
	// toJSONIdent()
	// fromJSON()
	vals := []string{"daniel"}
	modifyInplace(vals)
	fmt.Printf("vals: %v\n", vals)
}

func fetchIssues() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
