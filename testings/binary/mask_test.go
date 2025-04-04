package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mask(t *testing.T) {
	var rank byte = 7
	var suit byte = 3

	var card uint16 = uint16(rank)<<8 | uint16(suit)

	var outRank byte = byte(card >> 8)
	var outSuit byte = byte((card << 8) >> 8)

	assert.Equal(t, rank, outRank)
	assert.Equal(t, suit, outSuit)
}
