package iter

import (
	"testing"
)

func TestIteratorFromSlice(t *testing.T) {
	it := IntoIter([]any{1, "2", 3})
	hasNextValue(t, it, 1)
	hasNextValue(t, it, "2")
	hasNextValue(t, it, 3)
	noNextValue(t, it)
}
