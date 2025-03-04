package responder

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, data any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes()) // if fails to write no point in retry
}
