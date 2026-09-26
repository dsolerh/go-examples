package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/dsolerh/examples/concurrency_book/patterns/streamutils"
)

func main() {
	fmt.Println("Slow variant")
	example1()

	fmt.Println("\nFan-In variant")
	example2()
}

func randInt() interface{} { return rand.Intn(50000000) }

func example1() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := streamutils.ToInt(done, streamutils.RepeatFn(done, randInt))
	fmt.Println("Primes:")
	for prime := range streamutils.Take(done, streamutils.PrimeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}

func example2() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := streamutils.ToInt(done, streamutils.RepeatFn(done, randInt))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = streamutils.PrimeFinder(done, randIntStream)
	}

	for prime := range streamutils.Take(done, streamutils.FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
