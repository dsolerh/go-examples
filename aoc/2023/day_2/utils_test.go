package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_parseLine(t *testing.T) {
	tests := []struct {
		line []byte
		want gameInfo
	}{
		{
			[]byte("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
			gameInfo{id: 1, extractions: []extraction{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}},
		},
		{
			[]byte("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"),
			gameInfo{id: 2, extractions: []extraction{{blue: 1, green: 2}, {green: 3, blue: 4, red: 1}, {green: 1, blue: 1}}},
		},
		{
			[]byte("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"),
			gameInfo{id: 3, extractions: []extraction{{green: 8, blue: 6, red: 20}, {blue: 5, red: 4, green: 13}, {green: 5, red: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("line: '%s'", tt.line), func(t *testing.T) {
			if got := parseLine(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
