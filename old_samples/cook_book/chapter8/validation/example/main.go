package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dsolerh/examples/cook_book/chapter8/validation"
)

func main() {
	c := validation.New()
	http.HandleFunc("/", c.Process)
	fmt.Println("Listening on port :3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
