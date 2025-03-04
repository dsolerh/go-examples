package books

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var _ BookGetOne = (*mockBookGetOne)(nil)

type mockBookGetOne struct {
	mock.Mock
}

func (m *mockBookGetOne) GetBook(ctx context.Context, isbn string) (*bookContentData, error) {
	args := m.Called(ctx, isbn)

	return args.Get(0).(*bookContentData), args.Error(1)
}

type GetBookSuite struct {
	suite.Suite
}

func TestGetBookSuite(t *testing.T) {
	suite.Run(t, new(GetBookSuite))
}

func (s *GetBookSuite) TestGetBookFound() {
	isbn := "423443546"
	ctx := context.WithValue(
		context.Background(),
		chi.RouteCtxKey,
		&chi.Context{URLParams: chi.RouteParams{Keys: []string{"isbn"}, Values: []string{isbn}}},
	)
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/book/%s", isbn), nil)
	resp := httptest.NewRecorder()

	book := &bookContentData{
		Isbn:          isbn,
		Title:         "Super",
		Image:         "no image",
		Genre:         "drama",
		YearPublished: 2020,
	}
	repoMock := new(mockBookGetOne)
	repoMock.On("GetBook", req.Context(), isbn).Return(book, nil)

	// make the request
	GetBook(repoMock)(resp, req)

	s.Equal(http.StatusOK, resp.Code)

	body, err := io.ReadAll(resp.Body)
	s.NoError(err)

	bookJson, err := json.Marshal(book)
	s.NoError(err)
	s.JSONEq(string(bookJson), string(body))
}
