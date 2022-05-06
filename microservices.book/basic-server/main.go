package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	resp := helloResponse{Message: "Hello World!"}
	data, err := json.Marshal(resp)
	if err != nil {
		panic("Ooops")
	}
	fmt.Fprint(w, string(data))
}

func main() {
	port := 8080

	http.HandleFunc("/helloword", helloworldHandler)

	log.Printf("Server listening on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
