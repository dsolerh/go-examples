package main

func sum(vals ...int) int {
	var total int
	for _, v := range vals {
		total += v
	}
	return total
}

func max(vals ...int) int {
	ans := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > ans {
			ans = vals[i]
		}
	}
	return ans
}

func min(vals ...int) int {
	ans := vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] < ans {
			ans = vals[i]
		}
	}
	return ans
}
