package commonfunc

import "testing"

func TestTsConvert(t *testing.T) {
	testCases := []struct {
		ts   string
		from string
		to   string
		want string
	}{
		{
			ts:   "2021-03-08T19:12",
			from: "America/Los_Angeles",
			to:   "Asia/Jerusalem",
			want: "2021-03-09T05:12",
		},
	}
	for _, tC := range testCases {
		got, err := TsConvert(tC.ts, tC.from, tC.to)
		if err != nil {
			if err.Error() != tC.want {
				t.Errorf("got %q, want %q", err, tC.want)
			}
			continue
		}
		if got != tC.want {
			t.Errorf("got %q, want %q", got, tC.want)
		}
	}
}
