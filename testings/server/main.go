package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	fmt.Fprintf(
		w,
		"[%s] dbHost: %s | dbPort: %s | dbName: %s | dbUser: %s | dbPass: %s",
		time.Now(),
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPass,
	)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/healthz/db", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":8080", nil)
}
