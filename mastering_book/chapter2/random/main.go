package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	util "github.com/dsolerh/examples/mastering_book/chapter2/random/utils"
)

func main() {
	arguments := os.Args[1:]
	random := util.Random

	var (
		min   int = 0
		max   int = 100
		total int = 100
	)

	switch len(arguments) {
	case 5:
		fmt.Println("Too many arguments")
	case 4: // seed
		seed, err := strconv.Atoi(arguments[3])
		if err != nil {
			fmt.Printf("Invalid seed: %s", arguments[3])
			return
		}
		rand.Seed(int64(seed))
		fallthrough
	case 3: // total
		val, err := strconv.Atoi(arguments[2])
		if err != nil {
			fmt.Printf("Invalid total: %s", arguments[2])
			return
		}
		total = val
		fallthrough
	case 2: // max
		val, err := strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Printf("Invalid max: %s", arguments[1])
			return
		}
		max = val
		fallthrough
	case 1: // min
		val, err := strconv.Atoi(arguments[0])
		if err != nil {
			fmt.Printf("Invalid min: %s", arguments[0])
			return
		}
		min = val
		if min >= max {
			fmt.Printf("Invalid min(%d)>=max(%d)", min, max)
			return
		}
	case 0:
		fmt.Println("Using the default values!")
	}

	for i := 0; i < total; i++ {
		fmt.Print(random(min, max), " ")
	}
	fmt.Println()

}
