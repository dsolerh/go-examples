package books

import "context"

type BookGetOne interface {
	GetBook(ctx context.Context, isbn string) (*bookContentData, error)
}
