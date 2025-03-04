package books

import (
	"context"
	"database/sql"
)

type BooksRepository struct {
	db *sql.DB
}

func (r *BooksRepository) GetBook(ctx context.Context, isbn string) (*bookContentData, error) {
	row := r.db.QueryRowContext(ctx, "select isbn, title, image, genre, year_published from books where isbn = $1", isbn)

	bookData := bookContentData{}
	err := row.Scan(&bookData.Isbn, &bookData.Title, &bookData.Image, &bookData.Genre, &bookData.YearPublished)
	if err != nil {
		return nil, err
	}
	return &bookData, nil
}
