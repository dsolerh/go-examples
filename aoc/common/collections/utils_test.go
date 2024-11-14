package collections

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
	type args struct {
		s []byte
		n int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "take 1",
			args: args{
				s: []byte{'1'},
				n: 1,
			},
			want: []byte{'1'},
		},
		{
			name: "take 2",
			args: args{
				s: []byte{'1'},
				n: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Take(tt.args.s, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("take() got = %v, want %v", got, tt.want)
			}
		})
	}
}
