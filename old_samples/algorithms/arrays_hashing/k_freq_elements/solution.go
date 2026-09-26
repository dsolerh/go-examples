package kfreqelements

import "fmt"

func TopKElements(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums))

	for _, num := range nums {
		count[num]++
	}
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}
	fmt.Printf("freq: %v\n", freq)
	res := make([]int, 0, k)
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
