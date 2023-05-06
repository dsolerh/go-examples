package main

import (
	"fmt"
	"sort"
)

func main() {
	strArr := []string{"a", "ccc", "bb"}
	fmt.Println(strArr)
	sort.SliceStable(strArr, func(i, j int) bool {
		return len(strArr[i]) > len(strArr[j])
	})
	fmt.Println(strArr)

	sort.SliceStable(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})
	fmt.Println(strArr)
}
