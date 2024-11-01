package main

import (
	"bytes"
	"strconv"
)

type gameInfo struct {
	extractions []extraction
	id          int64
}

type extraction struct {
	red   int
	green int
	blue  int
}

var startIDPos = len([]byte("Game "))

func parseLine(line []byte) gameInfo {
	endIDPos := startIDPos + bytes.Index(line[startIDPos:], []byte(":"))
	id, _ := strconv.Atoi(string(line[startIDPos:endIDPos]))

	extractions := make([]extraction, 0)
	for _, b := range bytes.Split(line[endIDPos+1:], []byte(";")) {
		extraction := extraction{}
		for _, ext := range bytes.Split(b, []byte(",")) {
			data := bytes.Split(bytes.TrimSpace(ext), []byte(" "))
			color := data[1]
			amount, _ := strconv.Atoi(string(data[0]))

			if bytes.Equal(color, []byte("red")) {
				extraction.red += amount
			}
			if bytes.Equal(color, []byte("blue")) {
				extraction.blue += amount
			}
			if bytes.Equal(color, []byte("green")) {
				extraction.green += amount
			}
		}
		extractions = append(extractions, extraction)
	}

	return gameInfo{
		id:          int64(id),
		extractions: extractions,
	}
}
