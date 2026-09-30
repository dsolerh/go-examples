// Package httpapi exposes the auth service over HTTP/JSON.
package httpapi

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"microservices/auth/internal/service"
)

type Handler struct {
	auth *service.Auth
	log  *slog.Logger
}

func New(auth *service.Auth, log *slog.Logger) http.Handler {
	h := &Handler{auth: auth, log: log}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.health)
	mux.HandleFunc("POST /register", h.register)
	mux.HandleFunc("POST /login", h.login)
	mux.HandleFunc("POST /refresh", h.refresh)
	mux.HandleFunc("GET /validate", h.validate)
	return h.logRequests(mux)
}

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req credentials
	if !decode(w, r, &req) {
		return
	}
	id, err := h.auth.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		h.fail(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, id)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req credentials
	if !decode(w, r, &req) {
		return
	}
	t, err := h.auth.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		h.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if !decode(w, r, &req) {
		return
	}
	t, err := h.auth.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		h.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) validate(w http.ResponseWriter, r *http.Request) {
	raw, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !ok {
		h.fail(w, service.ErrUnauthorized)
		return
	}
	id, err := h.auth.Validate(raw)
	if err != nil {
		h.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, id)
}

func (h *Handler) fail(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		status = http.StatusBadRequest
	case errors.Is(err, service.ErrEmailTaken):
		status = http.StatusConflict
	case errors.Is(err, service.ErrInvalidCredentials), errors.Is(err, service.ErrUnauthorized):
		status = http.StatusUnauthorized
	default:
		h.log.Error("internal error", "err", err)
		err = errors.New("internal error")
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func (h *Handler) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.log.Info("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
