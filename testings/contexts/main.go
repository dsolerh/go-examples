package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	OuterContextTimeout = 8 * time.Millisecond
	InnerContextSleep   = 5 * time.Millisecond
	InnerContextWait    = 10 * time.Millisecond
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), OuterContextTimeout)

	addCancel(ctx, &wg)

	fmt.Printf("ctx: %v\n", ctx)
	fmt.Printf("ctx.Err(): %v\n", ctx.Err())

	cancel()
	wg.Wait()
}

func addCancel(ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go call(ctx, wg)
	time.Sleep(InnerContextSleep)
	cancel()
}

func call(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		fmt.Println("context canceled")
	case <-time.After(InnerContextWait):
		fmt.Println("waited for 10ms for the context to be canceled")
	}
	fmt.Printf("ctx: %v\n", ctx)
	fmt.Printf("ctx.Err(): %v\n", ctx.Err())
	wg.Done()
}
