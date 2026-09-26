package controllers

import (
	"encoding/json"
	"net/http"
)

// GetValue is a closure that wraps a HandlerFunc, if
// UseDefault is true value will always be "default" else it'll
// be whatever is stored in storage
func (c *Controller) GetValue(UseDefault bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		value := "default"
		if !UseDefault {
			value = c.storage.Get()
		}
		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}
	}
}
