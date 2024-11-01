package main

import (
	"cmp"
)

type T = struct{ A bool }

func main() {
	var t1 T
	var t2 T
	cmp.Or(t1, t2)
}
