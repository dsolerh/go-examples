package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_hasSymmetry(t *testing.T) {
	tests := []struct {
		patternLine []byte
		pos         int
		want        bool
	}{
		{
			patternLine: []byte("#.##..##."),
			pos:         0,
			want:        false,
		},
		{
			patternLine: []byte("#.##..##."),
			pos:         4,
			want:        true,
		},
		{
			patternLine: []byte("#.##..##."),
			pos:         6,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("pattern: %s is simetrical at %d ? %t", tt.patternLine, tt.pos, tt.want), func(t *testing.T) {
			if got := hasSymmetry(tt.patternLine, tt.pos); got != tt.want {
				t.Errorf("hasVerticalSimetry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verticalReflection(t *testing.T) {
	tests := []struct {
		pattern [][]byte
		want    int64
	}{
		{
			pattern: [][]byte{
				[]byte("#.##..##."),
				[]byte("..#.##.#."),
				[]byte("##......#"),
				[]byte("##......#"),
				[]byte("..#.##.#."),
				[]byte("..##..##."),
				[]byte("#.#.##.#."),
			},
			want: 5,
		},
		{
			pattern: [][]byte{
				[]byte("#...##..#"),
				[]byte("#....#..#"),
				[]byte("..##..###"),
				[]byte("#####.##."),
				[]byte("#####.##."),
				[]byte("..##..###"),
				[]byte("#....#..#"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(string(bytes.Join(tt.pattern, []byte(""))), func(t *testing.T) {
			if got := verticalReflection(tt.pattern); got != tt.want {
				t.Errorf("verticalReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_horizontalReflection(t *testing.T) {
	tests := []struct {
		pattern [][]byte
		want    int64
	}{
		{
			pattern: [][]byte{
				[]byte("#.##..##."),
				[]byte("..#.##.#."),
				[]byte("##......#"),
				[]byte("##......#"),
				[]byte("..#.##.#."),
				[]byte("..##..##."),
				[]byte("#.#.##.#."),
			},
			want: 0,
		},
		{
			pattern: [][]byte{
				[]byte("#...##..#"),
				[]byte("#....#..#"),
				[]byte("..##..###"),
				[]byte("#####.##."),
				[]byte("#####.##."),
				[]byte("..##..###"),
				[]byte("#....#..#"),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(string(bytes.Join(tt.pattern, []byte(""))), func(t *testing.T) {
			if got := horizontalReflection(tt.pattern); got != tt.want {
				t.Errorf("horizontalReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}
