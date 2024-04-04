package concurrentopperf

import (
	"math/rand"
	"sync"
)

const LONG_NUMBER = 10

func expensiveComputation() int {
	count := 0
	for i := 0; i < LONG_NUMBER; i++ {
		for j := 0; j < LONG_NUMBER; j++ {
			for k := 0; k < LONG_NUMBER; k++ {
				count += (i % 2) + (j % 5) - (k % 7)
			}
		}
	}
	return count
}

func concurrentFull(fn func() int, opCount int) int {
	resultChan := make(chan int, opCount)
	go func() {
		var wg sync.WaitGroup
		wg.Add(opCount)
		for i := 0; i < opCount; i++ {
			go func() {
				resultChan <- expensiveComputation()
				wg.Done()
			}()
		}
		wg.Wait()
		close(resultChan)
	}()

	total := 0
	for result := range resultChan {
		if result > total {
			total = result
		}
	}
	return total
}

func concurrentWithBuffer(fn func() int, opCount int, buffSize int) int {
	resultChan := make(chan int, buffSize)
	go func() {
		var wg sync.WaitGroup
		wg.Add(opCount)
		for i := 0; i < opCount; i++ {
			go func() {
				resultChan <- expensiveComputation()
				wg.Done()
			}()
		}
		wg.Wait()
		close(resultChan)
	}()

	total := 0
	for result := range resultChan {
		if result > total {
			total = result
		}
	}
	return total
}

func fd() {
	rand.ExpFloat64()
}
