package main

import (
	"aoc/data"
	"bufio"
	"flag"
	"fmt"
)

var Part = flag.Int("part", 1, "specify the excersice part")

const (
	MaxRed = iota + 12
	MaxGreen
	MaxBlue
)

func main() {
	flag.Parse()
	data, err := data.GetData(2023, 2)
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
