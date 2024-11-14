package api

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.ReadCloser, v interface{}) error {
	defer r.Close()
	return json.NewDecoder(r).Decode(v)
}
