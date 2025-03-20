package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooleanType_Schema(t *testing.T) {
	schema := new(BasicType).Boolean()
	wantSchema := BooleanTypeSchema{
		BasicTypeOptions: BasicTypeOptions{
			Type: BooleanTypeTag,
		},
	}
	is := assert.New(t)
	gotSchema := schema.Schema()
	is.Equal(wantSchema, gotSchema)
}
