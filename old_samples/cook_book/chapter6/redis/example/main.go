package main

import "github.com/dsolerh/examples/cook_book/chapter6/redis"

func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}
	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
