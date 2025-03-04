package e2e_tests

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func (s *E2ESuite) TestGetBookThatDoesNotExist() {
	client := http.Client{}

	resp, err := client.Get("http://localhost:8080/book/123456789")
	s.NoError(err)

	defer resp.Body.Close()

	s.Equal(http.StatusNotFound, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)

	s.JSONEq(`{"code": "001", "message": "No book with ISBN 123456789"}`, string(body))
}

func (s *E2ESuite) TestGetBookThatDoesExist() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	s.NoError(err)

	// provision the database
	isbn := "987654321"
	title := "There will be none"
	image := "https://some-domain/images/pic.jpg"
	genre := "Fiction"
	yearPublished := 2025
	res, err := db.Exec(
		"INSERT INTO books (isbn, title, image, genre, year_published) VALUES ($1, $2, $3, $4, $5)",
		isbn,
		title,
		image,
		genre,
		yearPublished)
	s.NoError(err)

	rows, err := res.RowsAffected()
	s.NoError(err)
	s.Equal(int64(1), rows)

	client := http.Client{}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/book/%s", isbn))
	s.NoError(err)

	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)

	expectedBody := fmt.Sprintf(`{
		"isbn": "%s",
		"title": "%s",
		"image": "%s",
		"genre": "%s",
		"year_published": %d
	}`, isbn, title, image, genre, yearPublished)
	s.JSONEq(expectedBody, string(body))
}
