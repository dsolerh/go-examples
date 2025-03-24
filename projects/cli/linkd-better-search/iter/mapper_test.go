package iter

import (
	"testing"
)

func TestMapper(t *testing.T) {
	it := IntoIter([]int{1, 2, 3}).Map(func(i int) int { return 2 * i })

	hasNextValue(t, it, 2)
	hasNextValue(t, it, 4)
	hasNextValue(t, it, 6)
	noNextValue(t, it)
}
