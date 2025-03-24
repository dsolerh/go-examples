package encodedecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAsciiDigit(t *testing.T) {
	is := assert.New(t)
	digits := []byte("0123456789")
	nonDigits := []byte(`qwertyuiop[]asdfghjkl;'\zxcvbnmm,/.`)

	for _, char := range digits {
		is.True(isAsciiDigit(char))
	}
	for _, char := range nonDigits {
		is.False(isAsciiDigit(char))
	}
}

func TestEncode(t *testing.T) {
	is := assert.New(t)
	strs := []string{"daniel", "soler", "hdez"}
	want := "6|daniel5|soler4|hdez"
	got := Encode(strs)
	is.Equal(want, got)
}

func TestDecode(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{
			str:  "6|daniel5|soler4|hdez",
			want: []string{"daniel", "soler", "hdez"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			is := assert.New(t)
			got, err := Decode(tc.str)
			is.NoError(err)
			is.Equal(tc.want, got)
		})
	}
}
