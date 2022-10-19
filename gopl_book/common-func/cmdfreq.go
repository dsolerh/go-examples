package commonfunc

import (
	"bufio"
	"os"
	"regexp"
)

/* Example history file
1264  traceroute google.com
1280  go env
1282  godoc
1283  go list
1284  go list ...
1286  cw go-ex.code-workspace
1443  mkdir gql-node-express-mongo
1444  code gql-node-express-mongo/
1519  cw go-ex.code-workspace
1521  cd go/
1523  cw go-ex.code-workspace
1686  cw gql-node-express-mongo.code-workspace
1729  go mod init examples/test
*/
var cmdRe = regexp.MustCompile(`[\d]{1,4}\s{2}go\s([a-z]+)`)

func CmdFreq(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	freqs := make(map[string]int)
	s := bufio.NewScanner(file)
	for s.Scan() {
		matches := cmdRe.FindStringSubmatch(s.Text())
		if len(matches) == 0 {
			continue
		}
		cmd := matches[1]
		freqs[cmd]++
	}

	if err = s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
