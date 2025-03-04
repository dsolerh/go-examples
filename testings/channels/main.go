package main

import (
	"fmt"
)

type empty = struct{}

func waitNil() <-chan empty {
	return nil
}

func waitImmediate() <-chan empty {
	ch := make(chan empty, 1)
	ch <- empty{}
	return ch
}

func read(ch <-chan empty) int {
	counter := 0
	for range ch {
		counter++
	}
	return counter
}

func main() {
	// select {
	// case <-waitNil():
	// 	fmt.Println("should never get here")
	// case <-waitImmediate():
	// 	fmt.Println("so perfect")
	// }

	// var wg sync.WaitGroup
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// ch := make(chan int, 0)
	// go func() {
	// 	ch <- 1
	// }()
	// fmt.Print(<-ch)

	arr1 := []int{1, 2, 3}
	arr2 := arr1[:1]
	arr2 = append(arr2, 4)
	arr2[0] = 10

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
}
