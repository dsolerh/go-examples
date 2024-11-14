package main

import (
	"bytes"
	"testing"
)

func TestGridIter_Col(t *testing.T) {
	tests := []struct {
		grid [][]byte
		col  int
		want []byte
	}{
		{
			grid: [][]byte{
				[]byte("#...##..#"),
				[]byte("#....#..#"),
				[]byte("..##..###"),
				[]byte("#####.##."),
				[]byte("#####.##."),
				[]byte("..##..###"),
				[]byte("#....#..#"),
			},
			col:  0,
			want: []byte("##.##.#"),
		},
	}
	for _, tt := range tests {
		t.Run(string(bytes.Join(tt.grid, []byte(""))), func(t *testing.T) {
			gIter := NewGridIter(tt.grid)
			elements := gIter.Col(tt.col)
			if !bytes.Equal(elements, tt.want) {
				t.Errorf("Collect() = %v, want %v", elements, tt.want)
			}
		})
	}
}

func TestGridIter_Row(t *testing.T) {
	tests := []struct {
		grid [][]byte
		col  int
		want []byte
	}{
		{
			grid: [][]byte{
				[]byte("#...##..#"),
				[]byte("#....#..#"),
				[]byte("..##..###"),
				[]byte("#####.##."),
				[]byte("#####.##."),
				[]byte("..##..###"),
				[]byte("#....#..#"),
			},
			col:  0,
			want: []byte("#...##..#"),
		},
	}
	for _, tt := range tests {
		t.Run(string(bytes.Join(tt.grid, []byte(""))), func(t *testing.T) {
			gIter := NewGridIter(tt.grid)
			elements := gIter.Row(tt.col)
			if !bytes.Equal(elements, tt.want) {
				t.Errorf("Collect() = %v, want %v", elements, tt.want)
			}
		})
	}
}
