package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	suscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clientRegistered sync.WaitGroup
	clientRegistered.Add(3)
	suscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clientRegistered.Done()
	})
	suscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!.")
		clientRegistered.Done()
	})
	suscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clientRegistered.Done()
	})
	button.Clicked.Broadcast()
	clientRegistered.Wait()
}
