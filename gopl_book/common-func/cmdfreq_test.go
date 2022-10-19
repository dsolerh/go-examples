package commonfunc

import (
	"reflect"
	"testing"
)

func TestCmdFreq(t *testing.T) {
	testCases := []struct {
		filename string
		want     map[string]int
	}{
		{
			filename: "cmd.hist.txt",
			want:     map[string]int{"env": 1, "list": 2, "mod": 3, "test": 4, "tool": 4},
		},
	}
	for _, tC := range testCases {
		got, err := CmdFreq(tC.filename)
		if err != nil {
			t.Errorf("got error: %q", err)
			continue
		}
		if !reflect.DeepEqual(got, tC.want) {
			t.Errorf("got %v, want %v", got, tC.want)
		}
	}
}
