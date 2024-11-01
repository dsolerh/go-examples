package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	records, err := readRecords()
	fmt.Printf("err: %v\n", err)
	fmt.Printf("records: %v\n", records)

	srecords, err := readStreamRecord()
	fmt.Printf("err: %v\n", err)
	fmt.Printf("srecords: %v\n", srecords)
}

type record []string

func (r record) validate() error {
	if len(r) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil
}

func fromLine(line string) record {
	return strings.Split(line, ",")
}

func readRecords() ([]record, error) {
	b, err := os.ReadFile("data.csv")
	if err != nil {
		return nil, err
	}

	content := string(b)                  // this converts from bytes to string
	lines := strings.Split(content, "\n") // split by line

	records := make([]record, 0)
	for i, line := range lines {
		// skip empty lines
		line = strings.TrimSpace(line) // remove white spaces from the beginning and the end of the line
		if line == "" {
			continue
		}
		record := fromLine(line)
		// validate the record
		if err := record.validate(); err != nil {
			return records, fmt.Errorf("entry at line [%d]'%s' was invalid: %w", i, line, err)
		}
		records = append(records, record)
	}
	return records, nil
}

func readStreamRecord() ([]record, error) {
	f, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	records := make([]record, 0)
	var lineNo int

	for scanner.Scan() {
		// remove white spaces from the beginning and the end of the line
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			lineNo++
			continue
		}

		record := fromLine(line)
		// validate the record
		if err := record.validate(); err != nil {
			return records, fmt.Errorf("entry at line [%d]'%s' was invalid: %w", lineNo, line, err)
		}
		records = append(records, record)
		lineNo++
	}

	return records, scanner.Err()
}
