package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"tdd_api/api/books"
	"tdd_api/internal/responder"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		responder.SendJSON(w, http.StatusOK, map[string]any{"status": "OK"})
	})

	router.Route("/book", books.RegisterRoutes(db))

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
