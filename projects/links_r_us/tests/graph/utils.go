package graphtest

import (
	"linkrus/internal/models/iterator"
	"sort"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertIteratedIDsMatch[T interface{ GetID() uuid.UUID }](assert *assert.Assertions, it iterator.Iterator[T], exp []uuid.UUID) {
	var got []uuid.UUID

	for it.Next() {
		got = append(got, it.Item().GetID())
	}

	// should not produce an error
	assert.Nil(it.Error())
	// should not produce an error
	assert.Nil(it.Close())

	sort.Slice(got, func(l, r int) bool { return got[l].String() < got[r].String() })
	sort.Slice(exp, func(l, r int) bool { return exp[l].String() < exp[r].String() })
	assert.Equal(got, exp)
}
