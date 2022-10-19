package main

import (
	"fmt"
	"net/http"

	"github.com/dsolerh/examples/cook_book/chapter8/controllers"
)

func main() {
	storage := controllers.MemStorage{}
	c := controllers.New(&storage)
	http.HandleFunc("/get", c.GetValue(false))
	http.HandleFunc("/get/default", c.GetValue(true))
	http.HandleFunc("/set", c.SetValue)
	fmt.Println("Listening on port :3333")
	panic(http.ListenAndServe(":3333", nil))
}
