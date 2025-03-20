package main

import "fmt"

func ff(i int) any {
	return i
}

func main() {
	var t any = int(1)
	f := ff(2)

	fmt.Printf("t: %v(%T)\n", t, t)
	fmt.Printf("f: %v(%T)\n", f, f)
}
