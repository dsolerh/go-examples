package main

import (
	"encoding/json"
	"net/http"
)

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloResponse{Message: "Hello " + name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
