package tests_test

import (
	"bytes"
	"cli-apps/w2p/replacer"
	"cli-apps/w2p/testdata"
	"fmt"
	"testing"
)

func Test_replaceSingleBlockBytes_files(t *testing.T) {
	testCases := testdata.Load(t, "import")

	r := replacer.NewImportReplacer(replacer.Replacements{
		// do/not/exist -> exist
		"do/not/exist": []byte("exist"),
	})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gotBytes, _, err := replacer.ReplaceSingleBlockBytes(bytes.NewReader(tc.FileContent), r)
			if err != nil {
				t.Errorf("replaceSingleBlockBytes() error = %v", err)
				return
			}
			if !bytes.Equal(gotBytes, tc.FileContentExpected) {
				t.Errorf("replaceSingleBlockBytes(%s) \ngot:\n%s\n\nwant:\n%s\ndiff:\n%s", name, gotBytes, tc.FileContentExpected, diffBytes(gotBytes, tc.FileContentExpected))
			}
		})
	}
}

func diffBytes(b1, b2 []byte) string {
	if len(b1) == 0 || len(b2) == 0 {
		return ""
	}

	n := min(len(b1), len(b2))
	for i := 0; i < n; i++ {
		if b1[i] != b2[i] {
			start := max(0, i-5)
			end := min(n, i+5)
			return fmt.Sprintf(
				"index: %d\n[b1] fragment: %s\n[b2] fragment: %s",
				i,
				b1[start:end], b2[start:end])
		}
	}
	if len(b1) != len(b2) {
		return "length is different"
	}
	return ""
}
