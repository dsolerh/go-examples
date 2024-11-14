package main

import (
	"aoc/data"
	"bufio"
	"flag"
	"fmt"
)

var Part = flag.Int("part", 2, "specify the excersice part")

func main() {
	flag.Parse()
	data, err := data.GetData(2023, 12)
	if err != nil {
		fmt.Printf("getting data err: %v\n", err)
		return
	}

	var solution int64
	if *Part == 1 {
		solution = solutionPart1(bufio.NewScanner(data))
	} else {
		solution = solutionPart2(bufio.NewScanner(data))
	}
	fmt.Printf("solution: %v\n", solution)
}
