package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2025, time.November, 25, 0, 0, 0, 0, time.UTC)
	newDate := date.Add(24 * 40 * time.Hour)
	fmt.Printf("newDate: %v\n", newDate)
}
