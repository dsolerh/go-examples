package testdata

import (
	"os"
	"path"
	"strings"
	"testing"
)

type TestData struct {
	FileContent         []byte
	FileContentExpected []byte
}

const dataDir = "./testdata"

func Load(t *testing.T, testname string) map[string]*TestData {
	t.Helper()
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		t.Fatalf("could not load the test case dir for test: %s | err: %v", testname, err)
		return nil
	}

	tc := make(map[string]*TestData, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryName := entry.Name()
		if strings.HasPrefix(entryName, testname) {
			name, found := strings.CutSuffix(entryName, ".expected")

			data, err := os.ReadFile(path.Join(dataDir, entryName))
			if err != nil {
				t.Fatalf("could not load the test case file: %s | err: %v", entryName, err)
				return nil
			}

			tdata := tc[name]
			if tdata == nil {
				tdata = &TestData{}
				tc[name] = tdata
			}

			if found {
				tdata.FileContentExpected = data
			} else {
				tdata.FileContent = data
			}
		}
	}

	return tc
}
