package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dsolerh/examples/cook_book/chapter8/handlers"
)

func main() {
	http.HandleFunc("/name", handlers.HelloHandler)
	http.HandleFunc("/greeting", handlers.GreetingHandler)
	fmt.Println("Listening on port :3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
