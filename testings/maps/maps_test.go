package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maps_go_string(t *testing.T) {
	var m1 = map[string]int64{
		"currency.coins": 12,
		"currency.gems":  13,
	}
	var m1Str1 = fmt.Sprintf("%v", m1)
	assert.Equal(t, `map[currency.coins:12 currency.gems:13]`, m1Str1)
	var m1Str2 = fmt.Sprintf("%#v", m1)
	assert.Equal(t, `map[string]int64{"currency.coins":12, "currency.gems":13}`, m1Str2)
	var m1Str3 = fmt.Sprintf("%+v", m1)
	assert.Equal(t, `map[currency.coins:12 currency.gems:13]`, m1Str3)
}
