package e2e_tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type E2ESuite struct {
	suite.Suite
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}

func (s *E2ESuite) TestHealthcheck() {
	client := http.Client{}

	resp, err := client.Get("http://localhost:8080/healthcheck")
	s.Nil(err, "no error should happen when caling the healthcheck")

	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	s.Nil(err, "no error should happen while reading the response body")

	s.JSONEq(`{"status": "OK"}`, string(body))
}
