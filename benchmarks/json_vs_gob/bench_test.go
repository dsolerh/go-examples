package jsonvsgob_test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"
)

type StructFullForTest struct {
	Str   string         `json:"str,omitempty"`
	Map   map[string]any `json:"map,omitempty"`
	Slice []any          `json:"slice,omitempty"`
	I64   int64          `json:"i_64,omitempty"`
	F64   float64        `json:"f_64,omitempty"`
	I     int            `json:"i,omitempty"`
	F32   float32        `json:"f_32,omitempty"`
	Bool  bool           `json:"bool,omitempty"`
	I8    int8           `json:"i_8,omitempty"`
}

type StructSimpleForTest struct {
	Str  string  `json:"str,omitempty"`
	I64  int64   `json:"i_64,omitempty"`
	F64  float64 `json:"f_64,omitempty"`
	I    int     `json:"i,omitempty"`
	F32  float32 `json:"f_32,omitempty"`
	Bool bool    `json:"bool,omitempty"`
	I8   int8    `json:"i_8,omitempty"`
}

var testStructFullData = StructFullForTest{
	I:    1234,
	I8:   32,
	I64:  18382732,
	F32:  2334.4342323,
	F64:  4234.42123121,
	Str:  "dasdvcbkjdl hfkjshakshds dshdiqwyeoej1902ey392 1ijp1dhkdjaoDHGOYAFG 98 2",
	Bool: false,
	Map: map[string]any{
		"dasdad":    1,
		" wdfsa":    213.32131,
		"dasddasad": "dsadsad",
		"das  sdad": false,
		"das d ad":  true,
	},
	Slice: []any{'c', "dsadasd", 1232, 34.342},
}

var testStructSimpleData = StructSimpleForTest{
	Str:  "asadasdsdasdadasdva a fsdfasvds vdv",
	I64:  31251,
	F64:  312,
	I:    3131,
	F32:  12354,
	Bool: false,
	I8:   23,
}

func init() {
	gob.Register(StructFullForTest{})
}

func BenchmarkEncode(b *testing.B) {
	var buff bytes.Buffer
	var encoders = []struct {
		encoder interface{ Encode(any) error }
		name    string
	}{
		{name: "gob", encoder: gob.NewEncoder(&buff)}, // NOTE: this is an expensive operation
		{name: "json", encoder: json.NewEncoder(&buff)},
	}
	for _, e := range encoders {
		b.Run(fmt.Sprintf("(%s) StructFull => ", e.name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := e.encoder.Encode(&testStructFullData)
				if err != nil {
					b.Fatalf("error while encoding: %v", err)
				}
				buff.Reset()
			}
		})
		b.Run(fmt.Sprintf("(%s) StructSimple => ", e.name), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := e.encoder.Encode(&testStructSimpleData)
				if err != nil {
					b.Fatalf("error while encoding: %v", err)
				}
				buff.Reset()
			}
		})
	}
}
