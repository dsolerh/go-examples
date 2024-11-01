package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := &scheduler{
		// executeAt: time.Now().Add(3 * time.Second),
		duration:  1 * time.Second,
		expiresIn: 4 * time.Second,
		immediate: true,
	}

	finish := BackgroundTaskManager(ctx, func() {
		fmt.Printf("now we are at: %s\n", time.Now())
		panic("just because I can")
	}, s)
	// wait for it to finish
	<-finish
}
