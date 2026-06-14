package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_contextCode(t *testing.T) {
	str, err := excecuteTemplate(contextCodeTemplate, contextCodeProps{
		PackageName: "test_package",
		Name:        "test_name",
		Props:       "Data int",
		SchemaId:    "test_schema_id",
	})
	assert.NoError(t, err)

	const expectedStr = `
package test_package
import "bitbucket.org/whatwapp/wakama/analytics"

const test_nameCtxSchemaId = "test_schema_id"
const test_nameCtxSchema = "depot:" + test_nameCtxSchemaId

type test_nameCtxProps struct {
	Data int
}

type test_nameCtx struct {
	analytics.Schema
	test_nameCtxProps
}

func (ctx *test_nameCtx) GetId() string {
	return test_nameCtxSchemaId
}

func Newtest_nameCtx(props test_nameCtxProps) *test_nameCtx {
	return &test_nameCtx{
		Schema:              analytics.Schema{Schema: test_nameCtxSchema},
		test_nameCtxProps: props,
	}
}`
	assert.Equal(t, expectedStr, str)
}
