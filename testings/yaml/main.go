package main

import (
	"log"

	"gopkg.in/yaml.v3"
)

var data_basic = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	example_basic()
}

func example_basic() {
	// Note: struct fields must be public in order for unmarshal to
	// correctly populate the data.
	type T struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}

	t := T{}

	err := yaml.Unmarshal([]byte(data_basic), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data_basic), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("--- m dump:\n%s\n\n", string(d))
}
