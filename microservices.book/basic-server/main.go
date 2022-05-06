package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type helloResponse struct {
	Message string `json:"message"`        // output: message
	Author  string `json:"-"`              // not display
	Date    string `json:"date,omitempty"` // not display if empty
	Id      int    `json:"id,string"`      // convert to string
}

type helloRequest struct {
	Name string `json:"name"`
}

// this is a faster implementation
func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	resp := helloResponse{
		Message: "Hello " + request.Name,
		Author:  "Daniel",
		Date:    time.Now().Format(time.Kitchen),
		Id:      123,
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(&resp)
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloworldHandler)

	log.Printf("Server listening on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
