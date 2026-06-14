package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/invopop/jsonschema"
	jsv "github.com/santhosh-tekuri/jsonschema/v6"
)

type Address struct {
	Street  string `json:"street"  jsonschema:"minLength=1"`
	City    string `json:"city"    jsonschema:"minLength=1"`
	ZipCode string `json:"zipCode" jsonschema:"pattern=^[0-9]{5}$"`
}

type Project struct {
	Name     string   `json:"name"             jsonschema:"minLength=1"`
	Priority int      `json:"priority"         jsonschema:"minimum=1,maximum=5"`
	Active   bool     `json:"active"`
	Skills   []string `json:"skills,omitempty" jsonschema:"uniqueItems=true,minItems=1"`
}

// Comment is self-referential: each comment may have a nested list of replies,
// which are themselves Comments. invopop emits Comment once under $defs and
// uses a $ref inside the "replies" items schema.
type Comment struct {
	ID      string    `json:"id"                jsonschema:"minLength=1"`
	Text    string    `json:"text"              jsonschema:"minLength=1,maxLength=500"`
	Replies []Comment `json:"replies,omitempty" jsonschema:"maxItems=50"`
}

// Category uses a pointer slice — handy when the recursion can be deep or
// when you want to distinguish "no children" from "empty children".
type Category struct {
	Name     string      `json:"name"               jsonschema:"minLength=1"`
	Children []*Category `json:"children,omitempty" jsonschema:"uniqueItems=false"`
}

type User struct {
	Name     string    `json:"name"               jsonschema:"minLength=1,maxLength=100"`
	Email    string    `json:"email"              jsonschema:"format=email"`
	Age      int       `json:"age"                jsonschema:"minimum=0,maximum=150"`
	Role     string    `json:"role"               jsonschema:"enum=admin,enum=user,enum=guest"`
	Tags     []string  `json:"tags,omitempty"     jsonschema:"uniqueItems=true"`
	Address  Address   `json:"address"`
	Projects []Project `json:"projects"           jsonschema:"minItems=1,maxItems=10"`
	Comments []Comment `json:"comments,omitempty"`
	Tree     *Category `json:"tree,omitempty"`
}

func main() {
	// 1. Generate a JSON Schema from the Go type.
	schema := jsonschema.Reflect(&User{})

	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatalf("marshal schema: %v", err)
	}
	fmt.Println("=== Generated JSON Schema ===")
	fmt.Println(string(schemaJSON))

	// 2. Compile the schema with santhosh-tekuri's validator.
	compiler := jsv.NewCompiler()
	schemaDoc, err := jsv.UnmarshalJSON(strings.NewReader(string(schemaJSON)))
	if err != nil {
		log.Fatalf("unmarshal schema: %v", err)
	}
	if err := compiler.AddResource("user.json", schemaDoc); err != nil {
		log.Fatalf("add resource: %v", err)
	}
	compiled, err := compiler.Compile("user.json")
	if err != nil {
		log.Fatalf("compile schema: %v", err)
	}

	// 3. Validate a couple of payloads.
	validPayload := `{
		"name": "Alice",
		"email": "alice@example.com",
		"age": 30,
		"role": "admin",
		"tags": ["go", "json"],
		"address": {"street": "Main 1", "city": "Madrid", "zipCode": "28001"},
		"projects": [
			{"name": "schema-gen", "priority": 1, "active": true,  "skills": ["go", "json-schema"]},
			{"name": "validator",  "priority": 3, "active": false, "skills": ["go"]}
		],
		"comments": [
			{
				"id": "c1", "text": "top-level",
				"replies": [
					{"id": "c1.1", "text": "first reply"},
					{
						"id": "c1.2", "text": "second reply",
						"replies": [
							{"id": "c1.2.1", "text": "deep reply"}
						]
					}
				]
			}
		],
		"tree": {
			"name": "root",
			"children": [
				{"name": "a", "children": [{"name": "a.1"}]},
				{"name": "b"}
			]
		}
	}`

	invalidPayload := `{
		"name": "",
		"email": "not-an-email",
		"age": -5,
		"role": "superuser",
		"tags": ["dup", "dup"],
		"address": {"street": "", "city": "Madrid", "zipCode": "ABC"},
		"projects": [
			{"name": "",       "priority": 9, "active": true,  "skills": ["dup", "dup"]},
			{"name": "legit",  "priority": 0, "active": "yes", "skills": []}
		],
		"comments": [
			{
				"id": "", "text": "",
				"replies": [
					{"id": "ok", "text": "fine", "replies": [
						{"id": "", "text": ""}
					]}
				]
			}
		],
		"tree": {
			"name": "",
			"children": [
				{"name": "", "children": [{"name": ""}]}
			]
		}
	}`

	fmt.Println("\n=== Validating valid payload ===")
	runValidation(compiled, validPayload)

	fmt.Println("\n=== Validating invalid payload ===")
	runValidation(compiled, invalidPayload)
}

func runValidation(sch *jsv.Schema, payload string) {
	var decoded any
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		log.Fatalf("unmarshal payload: %v", err)
	}
	err := sch.Validate(decoded)
	if err == nil {
		fmt.Println("OK")
		return
	}

	var vErr *jsv.ValidationError
	if !errors.As(err, &vErr) {
		fmt.Println("non-validation error:", err)
		return
	}

	fmt.Println("INVALID — flat list of leaf failures:")
	for _, leaf := range collectLeaves(vErr) {
		instance := "/" + strings.Join(leaf.InstanceLocation, "/")
		keyword := ""
		if leaf.ErrorKind != nil {
			keyword = "/" + strings.Join(leaf.ErrorKind.KeywordPath(), "/")
		}
		fmt.Printf("  at %-40s  kind=%-14T  keyword=%s\n", instance, leaf.ErrorKind, keyword)
	}

	fmt.Println("\nStructured (detailed) output:")
	out := vErr.DetailedOutput()
	j, _ := json.MarshalIndent(out, "  ", "  ")
	fmt.Println("  " + string(j))
}

// collectLeaves walks the Causes tree and returns only the leaves — those
// carry the actual failure kinds. Internal nodes are just grouping.
func collectLeaves(e *jsv.ValidationError) []*jsv.ValidationError {
	if len(e.Causes) == 0 {
		return []*jsv.ValidationError{e}
	}
	var out []*jsv.ValidationError
	for _, c := range e.Causes {
		out = append(out, collectLeaves(c)...)
	}
	return out
}
