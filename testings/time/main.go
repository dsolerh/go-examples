package main

import (
	"fmt"
	"time"
)

func main() {
	// testDuration(fmt.Sprintf("%dh", 24*5))
	// testDuration(fmt.Sprintf("%dh", 24*4))
	// maxD := time.Duration(math.MaxInt64)
	// fmt.Printf("maxD: %s\n", maxD)
	// now := time.Now()
	// fmt.Printf("now: %s\n", now)
	testTz()
}

func testTz() {
	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Printf("now: %s\n", now)
	fmt.Printf("now(UTC): %s\n", now.UTC())
	fmt.Printf("now(Roma): %s\n", now.UTC().In(loc))
	fmt.Printf("now(Roma)(UTC): %s\n", now.UTC().In(loc).UTC())
}

func test1() {
	str := time.Now().Format("2006-01-02T15:04:05")
	// 2006-01-02T15:04:05.999Z
	fmt.Printf("str: %v\n", str)
	isoStr := time.Now().Format("2006-01-02T15:04:05.999Z")
	fmt.Printf("isoStr: %v\n", isoStr)
}

func testDuration(durationStr string) {
	fmt.Printf("duration string: %s\n", durationStr)
	duration, err := time.ParseDuration(durationStr)
	fmt.Printf("error: %v\n", err)
	fmt.Printf("duration: %d\n", duration)
}
