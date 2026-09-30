package httpapi_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"microservices/auth/internal/httpapi"
	"microservices/auth/internal/service"
	"microservices/auth/internal/token"
	"microservices/auth/internal/user"
)

const (
	secret = "test-secret"
	email  = "ana@example.com"
)

func newServer(t *testing.T) *httptest.Server {
	t.Helper()
	return newServerWith(t, user.NewMemoryRepository())
}

func newServerWith(t *testing.T, repo user.Repository) *httptest.Server {
	t.Helper()
	tokens := token.NewManager(secret, time.Minute, time.Hour)
	auth := service.NewAuth(repo, tokens)
	srv := httptest.NewServer(httpapi.New(auth, slog.New(slog.DiscardHandler)))
	t.Cleanup(srv.Close)
	return srv
}

type response struct {
	code   int
	header http.Header
	body   map[string]string
}

// do sends a request. authHeader is the raw Authorization value ("" = none).
func do(t *testing.T, method, url, body, authHeader string) response {
	t.Helper()
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	out := map[string]string{}
	_ = json.Unmarshal(raw, &out)
	return response{code: res.StatusCode, header: res.Header, body: out}
}

func bearer(tok string) string { return "Bearer " + tok }

func refreshBody(tok string) string { return `{"refresh_token":"` + tok + `"}` }

// registerAndLogin creates the default user and returns its tokens.
func registerAndLogin(t *testing.T, srv *httptest.Server) (access, refresh string) {
	t.Helper()
	if r := do(t, "POST", srv.URL+"/register", `{"email":"ana@example.com","password":"supersecret"}`, ""); r.code != http.StatusCreated {
		t.Fatalf("register: got %d %v", r.code, r.body)
	}
	r := do(t, "POST", srv.URL+"/login", `{"email":"ana@example.com","password":"supersecret"}`, "")
	if r.code != http.StatusOK {
		t.Fatalf("login: got %d %v", r.code, r.body)
	}
	return r.body["access_token"], r.body["refresh_token"]
}

func TestFullFlow(t *testing.T) {
	srv := newServer(t)

	// Spaces and capitals, to check the email is normalized.
	r := do(t, "POST", srv.URL+"/register", `{"email":"  Ana@Example.com ","password":"supersecret"}`, "")
	if r.code != http.StatusCreated || r.body["email"] != email || r.body["user_id"] == "" {
		t.Fatalf("register: got %d %v", r.code, r.body)
	}
	if _, leaked := r.body["password"]; leaked {
		t.Fatal("register response must not contain the password")
	}
	userID := r.body["user_id"]

	if r := do(t, "POST", srv.URL+"/register", `{"email":"ana@example.com","password":"supersecret"}`, ""); r.code != http.StatusConflict {
		t.Fatalf("duplicate register: got %d", r.code)
	}

	// A different case, to check login is case-insensitive.
	r = do(t, "POST", srv.URL+"/login", `{"email":"ANA@example.com","password":"supersecret"}`, "")
	if r.code != http.StatusOK || r.body["access_token"] == "" || r.body["refresh_token"] == "" {
		t.Fatalf("login: got %d %v", r.code, r.body)
	}
	access, refresh := r.body["access_token"], r.body["refresh_token"]

	r = do(t, "GET", srv.URL+"/validate", "", bearer(access))
	if r.code != http.StatusOK || r.body["email"] != email || r.body["user_id"] != userID {
		t.Fatalf("validate: got %d %v", r.code, r.body)
	}

	r = do(t, "POST", srv.URL+"/refresh", refreshBody(refresh), "")
	if r.code != http.StatusOK || r.body["access_token"] == "" || r.body["refresh_token"] == "" {
		t.Fatalf("refresh: got %d %v", r.code, r.body)
	}

	// The refreshed access token must be usable.
	r = do(t, "GET", srv.URL+"/validate", "", bearer(r.body["access_token"]))
	if r.code != http.StatusOK || r.body["user_id"] != userID {
		t.Fatalf("validate refreshed token: got %d %v", r.code, r.body)
	}
}

func TestRegisterValidation(t *testing.T) {
	srv := newServer(t)
	cases := map[string]string{
		"invalid email":          `{"email":"not-an-email","password":"supersecret"}`,
		"empty email":            `{"email":"","password":"supersecret"}`,
		"password too short":     `{"email":"a@b.com","password":"1234567"}`,
		"empty password":         `{"email":"a@b.com","password":""}`,
		"missing fields":         `{}`,
		"password over 72 bytes": `{"email":"a@b.com","password":"` + strings.Repeat("x", 73) + `"}`,
		"not json":               `not json`,
		"empty body":             ``,
		"wrong field type":       `{"email":123,"password":"supersecret"}`,
		"body over 1MB limit":    `{"email":"a@b.com","password":"` + strings.Repeat("x", 1<<20) + `"}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			if r := do(t, "POST", srv.URL+"/register", body, ""); r.code != http.StatusBadRequest {
				t.Errorf("got %d, want 400", r.code)
			}
		})
	}

	for name, pw := range map[string]string{
		"password of exactly 8 bytes is accepted":  "12345678",
		"password of exactly 72 bytes is accepted": strings.Repeat("x", 72),
	} {
		t.Run(name, func(t *testing.T) {
			r := do(t, "POST", srv.URL+"/register", `{"email":"`+strings.Repeat("b", len(pw))+`@b.com","password":"`+pw+`"}`, "")
			if r.code != http.StatusCreated {
				t.Errorf("got %d, want 201", r.code)
			}
		})
	}
}

func TestLoginFailures(t *testing.T) {
	srv := newServer(t)
	registerAndLogin(t, srv)

	unknown := do(t, "POST", srv.URL+"/login", `{"email":"nobody@example.com","password":"supersecret"}`, "")
	wrongPass := do(t, "POST", srv.URL+"/login", `{"email":"ana@example.com","password":"wrong-pass"}`, "")

	for name, r := range map[string]response{"unknown email": unknown, "wrong password": wrongPass} {
		if r.code != http.StatusUnauthorized {
			t.Errorf("%s: got %d, want 401", name, r.code)
		}
	}
	// Same answer for both, so an attacker cannot find out which emails exist.
	if unknown.body["error"] != wrongPass.body["error"] {
		t.Errorf("error messages differ: %q vs %q", unknown.body["error"], wrongPass.body["error"])
	}

	if r := do(t, "POST", srv.URL+"/login", `not json`, ""); r.code != http.StatusBadRequest {
		t.Errorf("invalid json: got %d, want 400", r.code)
	}
}

func TestValidateRejects(t *testing.T) {
	srv := newServer(t)
	access, refresh := registerAndLogin(t, srv)

	expired, _ := token.NewManager(secret, -time.Minute, time.Hour).Issue("u1", email, token.Access)
	otherSecret, _ := token.NewManager("other-secret", time.Minute, time.Hour).Issue("u1", email, token.Access)

	cases := map[string]string{
		"no header":            "",
		"garbage token":        bearer("garbage"),
		"refresh token":        bearer(refresh),
		"expired token":        bearer(expired),
		"signed by other key":  bearer(otherSecret),
		"alg none":             bearer(forge(t, jwt.SigningMethodNone, jwt.UnsafeAllowNoneSignatureType, "auth-service", token.Access)),
		"wrong issuer":         bearer(forge(t, jwt.SigningMethodHS256, []byte(secret), "evil-service", token.Access)),
		"HS512 instead of 256": bearer(forge(t, jwt.SigningMethodHS512, []byte(secret), "auth-service", token.Access)),
		"tampered payload":     bearer(tamper(access)),
		"basic scheme":         "Basic " + access,
		"lowercase bearer":     "bearer " + access,
		"bearer without token": "Bearer ",
	}
	for name, header := range cases {
		t.Run(name, func(t *testing.T) {
			if r := do(t, "GET", srv.URL+"/validate", "", header); r.code != http.StatusUnauthorized {
				t.Errorf("got %d, want 401", r.code)
			}
		})
	}
}

func TestRefreshRejects(t *testing.T) {
	srv := newServer(t)
	access, _ := registerAndLogin(t, srv)

	expired, _ := token.NewManager(secret, time.Minute, -time.Minute).Issue("u1", email, token.Refresh)
	// Valid signature, but the user does not exist.
	ghost, _ := token.NewManager(secret, time.Minute, time.Hour).Issue("ghost-user", "ghost@example.com", token.Refresh)

	cases := map[string]string{
		"access token":  access,
		"expired token": expired,
		"unknown user":  ghost,
		"garbage":       "garbage",
		"empty":         "",
	}
	for name, tok := range cases {
		t.Run(name, func(t *testing.T) {
			if r := do(t, "POST", srv.URL+"/refresh", refreshBody(tok), ""); r.code != http.StatusUnauthorized {
				t.Errorf("got %d, want 401", r.code)
			}
		})
	}

	if r := do(t, "POST", srv.URL+"/refresh", `not json`, ""); r.code != http.StatusBadRequest {
		t.Errorf("invalid json: got %d, want 400", r.code)
	}
}

func TestInternalErrorsAreHidden(t *testing.T) {
	srv := newServerWith(t, failingRepo{})
	for _, path := range []string{"/register", "/login"} {
		r := do(t, "POST", srv.URL+path, `{"email":"ana@example.com","password":"supersecret"}`, "")
		if r.code != http.StatusInternalServerError {
			t.Errorf("%s: got %d, want 500", path, r.code)
		}
		if r.body["error"] != "internal error" {
			t.Errorf("%s: leaked error %q", path, r.body["error"])
		}
	}
}

func TestConcurrentRegisterSameEmail(t *testing.T) {
	srv := newServer(t)
	const n = 10
	codes := make(chan int, n)
	var wg sync.WaitGroup
	for range n {
		wg.Go(func() {
			codes <- do(t, "POST", srv.URL+"/register", `{"email":"ana@example.com","password":"supersecret"}`, "").code
		})
	}
	wg.Wait()
	close(codes)

	count := map[int]int{}
	for c := range codes {
		count[c]++
	}
	if count[http.StatusCreated] != 1 || count[http.StatusConflict] != n-1 {
		t.Fatalf("want 1x201 and %dx409, got %v", n-1, count)
	}
}

func TestRouting(t *testing.T) {
	srv := newServer(t)
	cases := []struct {
		method, path string
		want         int
	}{
		{"GET", "/healthz", http.StatusOK},
		{"GET", "/login", http.StatusMethodNotAllowed},
		{"GET", "/register", http.StatusMethodNotAllowed},
		{"POST", "/validate", http.StatusMethodNotAllowed},
		{"GET", "/does-not-exist", http.StatusNotFound},
	}
	for _, c := range cases {
		t.Run(c.method+" "+c.path, func(t *testing.T) {
			if r := do(t, c.method, srv.URL+c.path, "", ""); r.code != c.want {
				t.Errorf("got %d, want %d", r.code, c.want)
			}
		})
	}
}

func TestJSONResponses(t *testing.T) {
	srv := newServer(t)
	for name, r := range map[string]response{
		"success": do(t, "GET", srv.URL+"/healthz", "", ""),
		"error":   do(t, "GET", srv.URL+"/validate", "", ""),
	} {
		if ct := r.header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("%s: Content-Type = %q", name, ct)
		}
	}
}

// forge builds a token by hand, to test tokens our Manager would never make.
func forge(t *testing.T, method jwt.SigningMethod, key any, issuer string, kind token.Kind) string {
	t.Helper()
	claims := token.Claims{
		Email: email,
		Kind:  kind,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "u1",
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	s, err := jwt.NewWithClaims(method, claims).SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

// tamper changes one char in the payload, so the signature no longer matches.
func tamper(tok string) string {
	parts := strings.Split(tok, ".")
	p := []byte(parts[1])
	if p[0] == 'A' {
		p[0] = 'B'
	} else {
		p[0] = 'A'
	}
	parts[1] = string(p)
	return strings.Join(parts, ".")
}

type failingRepo struct{}

var errDB = errors.New("db down: connection refused at 10.0.0.5")

func (failingRepo) Create(context.Context, user.User) error            { return errDB }
func (failingRepo) ByEmail(context.Context, string) (user.User, error) { return user.User{}, errDB }
func (failingRepo) ByID(context.Context, string) (user.User, error)    { return user.User{}, errDB }
