package main

import (
	"encoding/json"
	"fmt"
)

const json_object = `
{
	"prop1": 1
}
`
const json_array = `
[1,2,3,4]
`

const json_number = `23`
const json_string = `"23"`
const json_bool = `true`
const json_null = `null`

func main() {
	for _, val := range []string{json_object, json_array, json_number, json_string, json_bool, json_null} {
		inspectJson([]byte(val))
	}
}

func inspectJson(json_data []byte) {
	var data any
	err := json.Unmarshal(json_data, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch data.(type) {
	case []any:
		fmt.Printf("this is an array %+v\n", data)

	case map[string]any:
		fmt.Printf("this is an object %+v\n", data)

	default:
		fmt.Printf("something else %T\n", data)
	}

}
