package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func Test_adjacentSymbol(t *testing.T) {
	schematic := [][]byte{
		{'.', '.', '.', '.'},
		{'.', '.', '.', '.'},
		{'.', '.', '*', '.'},
		{'.', '.', '.', '.'},
	}
	type args struct {
		row int
		col int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{0, 0}, false},
		{args{0, 1}, false},
		{args{0, 2}, false},
		{args{0, 3}, false},
		{args{1, 0}, false},
		{args{1, 1}, true},
		{args{1, 2}, true},
		{args{1, 3}, true},
		{args{2, 0}, false},
		{args{2, 1}, true},
		{args{2, 2}, false},
		{args{2, 3}, true},
		{args{3, 0}, false},
		{args{3, 1}, true},
		{args{3, 2}, true},
		{args{3, 3}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("row: %d col: %d has a symbol adjacent ? %t", tt.args.row, tt.args.col, tt.want), func(t *testing.T) {
			if got := adjacentSymbol(schematic, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("adjacentSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solutionPart1(t *testing.T) {
	const expectedSolution = 4361
	solution := solutionPart1(bufio.NewScanner(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)))

	if solution != expectedSolution {
		t.Errorf("the solution shouldbe %d but got %d", expectedSolution, solution)
	}
}

func Test_solutionPart1_edges(t *testing.T) {
	const expectedSolution = 0
	solution := solutionPart1(bufio.NewScanner(strings.NewReader(`123.422.123`)))

	if solution != expectedSolution {
		t.Errorf("the solution shouldbe %d but got %d", expectedSolution, solution)
	}
}
