package commonfunc

import "testing"

func TestFilter(t *testing.T) {
	testCases := []struct {
		values []int
		fn     func(int) bool
		want   []int
	}{
		{
			values: []int{1, 2, 3, 4},
			fn:     func(i int) bool { return i%2 == 0 },
			want:   []int{2, 4},
		},
	}
	for _, tC := range testCases {
		got := Filter(tC.fn, tC.values)

		if len(tC.want) != len(got) {
			t.Errorf("got %v, want %v", got, tC.want)
		}
	}
}
