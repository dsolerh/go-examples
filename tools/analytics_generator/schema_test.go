package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Schema(t *testing.T) {
	const schemaStr = `{
		"self": {
			"namespace": "com.whatwapp.euchre",
			"category": "analytics",
			"name": "transaction_ctx",
			"version": "1-0-0"
		},
		"type": "object",
		"properties": {
			"id": {
				"type": "string"
			},
			"coins": {
				"type": "integer"
			},
			"league.points": {
				"type": "integer"
			},
			"xp.points": {
				"type": "integer"
			}
		},
		"required": [],
		"meta": {
			"vertical": true,
			"dynamic": "integer"
		},
		"$schema": "https://schema.whatwapp.io/api/v1/schemas/com.whatwapp/self-schema/schema/1-0-0",
		"$id": "depot:com.whatwapp.euchre/analytics/transaction_ctx/1-0-0"
	}`

	schema, err := Json[Schema](schemaStr)
	assert.NoError(t, err)
	assert.Equal(t, Schema{
		SchemaType: SchemaType{
			Type: TObject,
			Properties: TProperties{
				"id":            SchemaType{Type: TString},
				"coins":         SchemaType{Type: TInteger},
				"league.points": SchemaType{Type: TInteger},
				"xp.points":     SchemaType{Type: TInteger},
			},
		},
		SchemaMeta: SchemaMeta{
			Self: SelfInfo{
				Namespace: "com.whatwapp.euchre",
				Category:  "analytics",
				Name:      "transaction_ctx",
				Version:   "1-0-0",
			},
			Meta: MetaInfo{
				Vertical: true,
				Dynamic:  "integer",
			},
			Schema: "https://schema.whatwapp.io/api/v1/schemas/com.whatwapp/self-schema/schema/1-0-0",
			Id:     "depot:com.whatwapp.euchre/analytics/transaction_ctx/1-0-0",
		},
	}, schema)
}
