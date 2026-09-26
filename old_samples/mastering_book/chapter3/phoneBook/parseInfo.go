package main

import (
	"fmt"
	"regexp"
	"time"
)

const ISOlayout = "2006-01-02T15:04:05Z"

func matchSurName(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func parseRecord(fields []string) (*Record, error) {
	if len(fields) != 4 {
		return nil, fmt.Errorf("invalid number of records: %d", len(fields))
	}
	if !matchSurName(fields[0]) {
		return nil, fmt.Errorf("invalid name: %s", fields[0])
	}
	if !matchSurName(fields[1]) {
		return nil, fmt.Errorf("invalid surname: %s", fields[1])
	}
	if !matchInt(fields[2]) {
		return nil, fmt.Errorf("invalid phone: %s", fields[2])
	}
	last, err := time.Parse(ISOlayout, fields[3])
	if err != nil {
		return nil, err
	}

	return &Record{
		Name:       fields[0],
		Surname:    fields[1],
		Number:     fields[2],
		LastAccess: last,
	}, nil
}
