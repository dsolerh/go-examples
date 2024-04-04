package concurrentopperf

import (
	"fmt"
	"testing"
)

func Benchmark_expensiveComputation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expensiveComputation()
	}
}

func Benchmark_concurrentFull(b *testing.B) {
	var inputs = []int{
		100, 1000, 10000,
	}
	for _, amount := range inputs {
		b.Run(fmt.Sprintf("with amount: %d", amount), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				concurrentFull(expensiveComputation, amount)
			}
		})
	}
}

func Benchmark_concurrentWithBuffer(b *testing.B) {
	var inputs = [][2]int{
		{100, 20},
		{1000, 20},
		{10000, 20},
	}
	for _, input := range inputs {
		b.Run(fmt.Sprintf("with input: %v", input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				concurrentWithBuffer(expensiveComputation, input[0], input[1])
			}
		})
	}
}
