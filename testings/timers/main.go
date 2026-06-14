package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	t := time.AfterFunc(1*time.Hour, func() {
		counter++
	})
	wasRescheduled := t.Reset(-1)
	fmt.Printf("wasRescheduled: %v\n", wasRescheduled)
	wasStoped := t.Stop()
	if wasStoped {
		fmt.Printf("the timer was stoped, counter = %d\n", counter)
	} else {
		fmt.Printf("the timer was not stoped, counter = %d\n", counter)
		select {
		case t := <-t.C:
			fmt.Printf("there was something in the C chan '%v'?\n", t)
		default:
			fmt.Println("there was nothing in the C chan")
		}
	}

	ch := make(chan struct{}, 1)
	t = time.AfterFunc(1*time.Nanosecond, func() {
		counter++
		ch <- struct{}{}
	})
	<-ch
	wasStoped = t.Stop()
	if wasStoped {
		fmt.Printf("the timer was stoped, counter = %d\n", counter)
	} else {
		fmt.Printf("the timer was not stoped, counter = %d\n", counter)
		select {
		case t := <-t.C:
			fmt.Printf("there was something in the C chan '%v'?\n", t)
		default:
			fmt.Println("there was nothing in the C chan")
		}
	}
}
