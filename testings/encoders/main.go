package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math"
)

type StructSimpleForTest struct {
	Str  string  `json:"str,omitempty"`
	I64  int64   `json:"i_64,omitempty"`
	F64  float64 `json:"f_64,omitempty"`
	I    int     `json:"i,omitempty"`
	F32  float32 `json:"f_32,omitempty"`
	Bool bool    `json:"bool,omitempty"`
	I8   int8    `json:"i_8,omitempty"`
}

var simpleData = StructSimpleForTest{
	Str:  "1234",
	I64:  math.MaxInt64,
	F64:  math.MaxFloat64,
	I:    math.MaxInt,
	F32:  math.MaxFloat32,
	Bool: false,
	I8:   127,
}

type StructNestedForTest struct {
	Nested struct {
		Str   string `json:"str,omitempty"`
		Slice []struct {
			Str string  `json:"str,omitempty"`
			I64 int64   `json:"i_64,omitempty"`
			F64 float64 `json:"f_64,omitempty"`
		} `json:"slice,omitempty"`
		I64 int64   `json:"i_64,omitempty"`
		F64 float64 `json:"f_64,omitempty"`
	} `json:"nested,omitempty"`
	I    int     `json:"i,omitempty"`
	F32  float32 `json:"f_32,omitempty"`
	Bool bool    `json:"bool,omitempty"`
	I8   int8    `json:"i_8,omitempty"`
}

var nestedData = StructNestedForTest{
	Nested: struct {
		Str   string "json:\"str,omitempty\""
		Slice []struct {
			Str string  "json:\"str,omitempty\""
			I64 int64   "json:\"i_64,omitempty\""
			F64 float64 "json:\"f_64,omitempty\""
		} "json:\"slice,omitempty\""
		I64 int64   "json:\"i_64,omitempty\""
		F64 float64 "json:\"f_64,omitempty\""
	}{
		Str: "",
		Slice: []struct {
			Str string  "json:\"str,omitempty\""
			I64 int64   "json:\"i_64,omitempty\""
			F64 float64 "json:\"f_64,omitempty\""
		}{
			{
				Str: "",
				I64: 0,
				F64: 0,
			},
			{
				Str: "",
				I64: 0,
				F64: 0,
			},
			{
				Str: "",
				I64: 0,
				F64: 0,
			},
		},
		I64: 0,
		F64: 0,
	},
	I:    0,
	F32:  0,
	Bool: false,
	I8:   0,
}

func main() {
	checkSize(&simpleData)
	fmt.Println()
	checkSize(&nestedData)
}

func checkSize[D any](data D) {
	var buff bytes.Buffer
	var err error
	err = gob.NewEncoder(&buff).Encode(&data)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("buff.String(): %v\n", buff.String())
	fmt.Printf("gob bytes: %v\n", len(buff.Bytes()))

	buff.Reset()
	err = json.NewEncoder(&buff).Encode(&data)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("buff.String(): %v\n", buff.String())
	fmt.Printf("json bytes: %v\n", len(buff.Bytes()))
}
