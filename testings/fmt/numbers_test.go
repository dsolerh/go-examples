package fmtest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numbers(t *testing.T) {
	assert.Equal(t, "v: 001", fmt.Sprintf("v: %03d", 1))
}
