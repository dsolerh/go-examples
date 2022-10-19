package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSVFile(fpath string) error {
	fileInfo, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s not a regular file", fpath)
	}
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		record, err := parseRecord(line)
		if err != nil {
			return err
		}
		data = append(data, record)
		// fmt.Println(*record)
	}
	return nil
}

func saveCSVFile(fpath string) error {
	csvfile, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess.Format(ISOlayout)}
		err := csvwriter.Write(temp)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
	csvwriter.Flush()
	return nil
}
