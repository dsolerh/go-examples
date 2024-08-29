package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

type Arr []int

func (a Arr) Sort() {
	slices.Sort(a)
}

func main() {
	// strArr := []string{"a", "ccc", "bb"}
	// fmt.Println(strArr)
	// sort.SliceStable(strArr, func(i, j int) bool {
	// 	return len(strArr[i]) > len(strArr[j])
	// })
	// fmt.Println(strArr)

	// sort.SliceStable(strArr, func(i, j int) bool {
	// 	return strArr[i] < strArr[j]
	// })
	// fmt.Println(strArr)
	// a := Arr{4, 2, 1, 3}
	// a.Sort()
	// fmt.Printf("a: %v\n", a)
	// src := 1
	// dst := 0
	// str := "sdadad"
	// fmt.Printf("src: %016b\ndst: %016b\nenc: %064b\n", src, dst, src<<8+dst<<2)

	// aa := func () {
	// 	ss := struct{a int}{}
	// 	return &ss
	// }()

	// os.Getenv("")
	// go func(){}()
	animationTime := 950 * time.Millisecond
	ticks := math.Round(animationTime.Seconds() * 8.0)
	fmt.Printf("ticks: %v\n", ticks)
}
