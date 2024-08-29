package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type StructLog struct{}

func (l *StructLog) ParseLine(line []byte) string {
	var obj map[string]any
	err := json.Unmarshal(line, &obj)
	if err != nil {
		return fmt.Sprintf("err: %v | raw_line: %s\n", err, line)
	} else {
		return fmt.Sprintf("obj: %v\n", obj)
	}
}

func main() {
	output := os.Stdout
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	l := &StructLog{}
	for counter := 0; scanner.Scan(); counter++ {
		line := scanner.Bytes()
		fmt.Fprintf(output, "%d | %s", counter, l.ParseLine(line))
	}
}
