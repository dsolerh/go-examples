package main

import (
	"fmt"
	"time"
)

const MIN_TIMEOUT_TICKS = 40

func getTimeoutForRemainingTicks(remainingTicks int, tickRate int) time.Duration {
	// at least MIN_TIMEOUT_TICKS ticks of time
	// and substract a small amount of time to try to have the timeout
	// before the data is needed. 10ms seems fine.
	ticks := max(remainingTicks, MIN_TIMEOUT_TICKS)
	return time.Duration(ticks*1000/tickRate)*time.Millisecond - 10*time.Millisecond
}

func main() {
	timeout := getTimeoutForRemainingTicks(0, 8)
	fmt.Printf("timeout: %s\n", timeout)
}
