package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, data any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes()) //nolint:errcheck // if fails to write no point in retry
}
