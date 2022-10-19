package main

import "github.com/dsolerh/examples/cook_book/chapter6/database"

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}
	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
