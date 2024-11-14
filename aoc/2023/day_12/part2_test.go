package main

import "testing"

func Test_iterative(t *testing.T) {
	for _, tt := range cases {
		t.Run(string(tt.springs), func(t *testing.T) {
			if got := iterative(tt.springs, tt.damagedSprings); got != tt.want {
				t.Errorf("iterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
