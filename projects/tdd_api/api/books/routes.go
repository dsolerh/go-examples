package books

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	commondata "tdd_api/internal/common_data"
	"tdd_api/internal/responder"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(db *sql.DB) func(router chi.Router) {
	repo := &BooksRepository{db}
	return func(router chi.Router) {
		router.Get("/{isbn}", GetBook(repo))
	}
}

func GetBook(repo BookGetOne) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isbn := chi.URLParam(r, "isbn")

		bookData, err := repo.GetBook(r.Context(), isbn)
		if errors.Is(err, sql.ErrNoRows) {
			responder.SendJSON(w, http.StatusNotFound, commondata.JsonError{Code: "001", Message: fmt.Sprintf("No book with ISBN %s", isbn)})
			return
		} else if err != nil {
			responder.SendJSON(w, http.StatusInternalServerError, commondata.JsonError{Code: "002", Message: "An error occurred while retrieving the book"})
			return
		}
		responder.SendJSON(w, http.StatusOK, bookData)
	}
}
