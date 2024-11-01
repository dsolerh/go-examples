package searcher

import (
	"aoc/common/digits"
	"bytes"
	"slices"
	"testing"
)

func Test_searcher_Next(t *testing.T) {
	type fields struct {
		line      []byte
		searchers []SearchFn
	}
	tests := []struct {
		name      string
		fields    fields
		want      bool
		wantFound []byte
		wantAt    int
		wantIndex int
	}{
		{
			name: "",
			fields: fields{
				line: []byte(`12ad43ad22`),
				searchers: []func(data []byte) []byte{func(data []byte) []byte {
					if digits.IsDigit(data[0]) {
						return []byte{data[0]}
					}
					return nil
				}},
			},
			want:      true,
			wantFound: []byte("1"),
			wantAt:    0,
			wantIndex: 1,
		},
		{
			name: "",
			fields: fields{
				line: []byte(`12ad43ad22`),
				searchers: []func(data []byte) []byte{func(data []byte) []byte {
					if !digits.IsDigit(data[0]) {
						return []byte{data[0]}
					}
					return nil
				}},
			},
			want:      true,
			wantFound: []byte("a"),
			wantAt:    2,
			wantIndex: 3,
		},
		{
			name: "",
			fields: fields{
				line: []byte(`12ad43ad22`),
				searchers: []func(data []byte) []byte{func(data []byte) []byte {
					if bytes.HasPrefix(data, []byte("12")) {
						return []byte("12")
					}
					return nil
				}},
			},
			want:      true,
			wantFound: []byte("12"),
			wantAt:    0,
			wantIndex: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &lineSearcher{
				line:      tt.fields.line,
				searchers: tt.fields.searchers,
			}
			if got := s.Next(); got != tt.want {
				t.Errorf("searcher.next() = %t, want %t", got, tt.want)
			}
			if !bytes.Equal(s.found, tt.wantFound) {
				t.Errorf("searcher.found = %s, want %s", s.found, tt.wantFound)
			}
			if s.at != tt.wantAt {
				t.Errorf("searcher.at = %d, want %d", s.at, tt.wantAt)
			}
			if s.currentIndex != tt.wantIndex {
				t.Errorf("searcher.currentIndex = %d, want %d", s.currentIndex, tt.wantIndex)
			}
		})
	}
}

func Test_searcher_Numbers(t *testing.T) {
	s := &lineSearcher{
		line:      []byte("....232.633.....803.."),
		searchers: []func(data []byte) []byte{Numbers},
	}

	type item struct {
		data []byte
		at   int
	}
	items := make([]item, 0)

	for s.Next() {
		data, at := s.Item()
		items = append(items, item{data, at})
	}

	expectedItems := []item{
		{[]byte("232"), 4},
		{[]byte("633"), 8},
		{[]byte("803"), 16},
	}
	equal := func(i1, i2 item) int {
		if i1.at != i2.at {
			return 1
		}
		if !bytes.Equal(i1.data, i2.data) {
			return 1
		}
		return 0
	}

	if slices.CompareFunc(items, expectedItems, equal) != 0 {
		t.Errorf("expected: %v but got: %v", expectedItems, items)
	}
}
