package bytes

import "fmt"

func ExampleTag() {
	data := []byte("(abc)")
	parser := Tag([]byte("(ab"))
	found, remaining := parser(data)
	fmt.Printf("found: %s\n", found)
	fmt.Printf("remaining: %s\n", remaining)
	// Output:
	// found: (ab
	// remaining: c)
}
