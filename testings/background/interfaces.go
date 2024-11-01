package main

import "time"

/*
Idea: build a function that executes background tasks:
	- needs to receive a function which is going to execute with a given schedule
	- for the schedule:
		- will receive an interface that will implement a method that will return from a channel when
		  the function should execute and another that will return from a channel if the background task should be stopped.
		- this will enable to schedule with flexibility like:
			- execute every x minutes
			- execute every x minutes for y minutes (similar behaviour can be achieved with the context)
			- execute at timestamp x
			- note: it's not intended to be used like `cron` by specifying a '* * * * *' schedule format
	- needs to return a channel that will be closed if the background tasks are stoped
	- needs to receive a context that will control it's execution (stopping)
	- the functions that will receive will not receive params or return values
*/

type empty = struct{}

// Scheduler represents someting that can create a schedule.
type Scheduler interface {
	New() schedule
}

// schedule abstracts away the implementation for a schedule
type schedule interface {
	close()
	ready() <-chan time.Time
	done() <-chan time.Time
}
