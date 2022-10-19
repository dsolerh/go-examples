package main

import (
	"fmt"
	"time"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess time.Time
}

var data []*Record
var index map[string]int

func createIndex() {
	index = make(map[string]int)
	for i := 0; i < len(data); i++ {
		index[data[i].Number] = i
	}
}

func insertRecord(r *Record) error {
	data = append(data, r)
	index[r.Number] = len(data) - 1
	return nil
}

func deleteRecord(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("no record found with key: %s", key)
	}
	data = append(data[:i], data[i+1:]...)
	return nil
}

func searchRecord(number string) *Record {
	i, ok := index[number]
	if !ok {
		return nil
	}
	return data[i]
}

func listRecord() {
	fmt.Println("#\tName\t\tSurname\t\tNumber")
	for i, v := range data {
		fmt.Printf("%d\t%s\t\t%s\t\t%s\n", i, v.Name, v.Surname, v.Number)
	}
}
