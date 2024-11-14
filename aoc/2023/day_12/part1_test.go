package main

import (
	"testing"
)

func Test_recur(t *testing.T) {
	for _, tt := range cases {
		t.Run(string(tt.springs), func(t *testing.T) {
			if got := recursive(tt.springs, tt.damagedSprings); got != tt.want {
				t.Errorf("recur() = %v, want %v", got, tt.want)
			}
		})
	}
}
