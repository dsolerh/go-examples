package commonfunc

func Filter(pred func(int) bool, values []int) []int {
	var filtered []int
	for _, v := range values {
		if pred(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
