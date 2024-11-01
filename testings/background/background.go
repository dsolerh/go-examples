package main

import (
	"context"
	"fmt"
	"time"
)

// BackgroundTaskManager uses a Scheduler to run a func fn. It will return a channel that will be
// closed upon completion of the tasks.
func BackgroundTaskManager(ctx context.Context, fn func(), scheduler Scheduler) <-chan empty {
	done := make(chan empty)

	go func() {
		schedule := scheduler.New()
		// will close the scheduler
		defer schedule.close()
		defer close(done)

		fmt.Printf("Starting job loop at %s\n", time.Now())
		for {
			select {
			// if the schedule is ready execute the function
			case t := <-schedule.ready():
				func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Printf("operation at %s paniqued: %v\n", t, err)
						}
					}()
					fn()
				}()

			// if the schedule is done return
			case <-schedule.done():
				return

			// if the context is done also return
			case <-ctx.Done():
				return
			}
		}
	}()
	return done
}
