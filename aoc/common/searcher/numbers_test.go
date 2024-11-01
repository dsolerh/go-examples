package searcher

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNumbers(t *testing.T) {
	tests := []struct {
		data []byte
		want []byte
	}{
		{
			data: []byte("1...."),
			want: []byte("1"),
		},
		{
			data: []byte("123...."),
			want: []byte("123"),
		},
		{
			data: []byte("123456"),
			want: []byte("123456"),
		},
		{
			data: []byte("...456"),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("data: %s -> want: %s", tt.data, tt.want), func(t *testing.T) {
			if got := Numbers(tt.data); !bytes.Equal(got, tt.want) {
				t.Errorf("Numbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
