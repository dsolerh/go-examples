package main

import (
	"fmt"
	"testing"
)

func Test_isDigit(t *testing.T) {
	tests := []struct {
		b    byte
		want bool
	}{
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
		//
		{'a', false},
		{'b', false},
		{'c', false},
		{'A', false},
		{'B', false},
		{'C', false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("'%c' is digit => %t", tt.b, tt.want), func(t *testing.T) {
			if got := isDigit(tt.b); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDigit(t *testing.T) {
	tests := []struct {
		b    byte
		want int64
	}{
		{'0', 0},
		{'1', 1},
		{'2', 2},
		{'3', 3},
		{'4', 4},
		{'5', 5},
		{'6', 6},
		{'7', 7},
		{'8', 8},
		{'9', 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("'%c' -> %d", tt.b, tt.want), func(t *testing.T) {
			if got := toDigit(tt.b); got != tt.want {
				t.Errorf("toDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeNumber(t *testing.T) {
	tests := []struct {
		line []byte
		want int64
	}{
		{[]byte("1abc2"), 12},
		{[]byte("pqr3stu8vwx"), 38},
		{[]byte("a1b2c3d4e5f"), 15},
		{[]byte("treb7uchet"), 77},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("line: '%s' -> num: %d", tt.line, tt.want), func(t *testing.T) {
			if got := computeNumberV1(tt.line); got != tt.want {
				t.Errorf("computeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstNumber(t *testing.T) {
	tests := []struct {
		line []byte
		want int64
	}{
		{[]byte("two1nine"), 2},
		{[]byte("eightwothree"), 8},
		{[]byte("abcone2threexyz"), 1},
		{[]byte("xtwone3four"), 2},
		{[]byte("4nineeightseven2"), 4},
		{[]byte("zoneight234"), 1},
		{[]byte("7pqrstsixteen"), 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("line: '%s' -> first: %d", tt.line, tt.want), func(t *testing.T) {
			if got := firstNumber(tt.line); got != tt.want {
				t.Errorf("firstNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastNumber(t *testing.T) {
	tests := []struct {
		line []byte
		want int64
	}{
		{[]byte("two1nine"), 9},
		{[]byte("eightwothree"), 3},
		{[]byte("abcone2threexyz"), 3},
		{[]byte("xtwone3four"), 4},
		{[]byte("4nineeightseven2"), 2},
		{[]byte("zoneight234"), 4},
		{[]byte("7pqrstsixteen"), 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("line: '%s' -> last: %d", tt.line, tt.want), func(t *testing.T) {
			if got := lastNumber(tt.line); got != tt.want {
				t.Errorf("lastNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
