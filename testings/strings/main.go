package main

import (
	"fmt"
	"time"
)

func main() {
	// names := []string{"daniel", "  daniel  ", "%%%%daniel%%%", " % %%% daniel"}
	// for _, name := range names {
	// 	trimName := string(bytes.TrimLeft([]byte(name), "% "))
	// 	fmt.Printf("name: [%v]\n", name)
	// 	fmt.Printf("trimName: [%v]\n", trimName)
	// 	fmt.Println()
	// }
	a := []byte{91, 34, 34, 44, 49, 93}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a: %s\n", a)
	time.Now()
}
