package nkclient

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// APIError is returned when the Nakama server responds with a non-2xx status.
// It exposes the HTTP status plus the structured error body so callers can
// branch on either.
type APIError struct {
	Status  int
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("nkclient: api error: status=%d code=%d: %s", e.Status, e.Code, e.Message)
}

// AuthenticateCustom signs in (or signs up) using a custom identifier.
//
// customID is the caller-supplied stable identifier for the player. When
// req.Create is true the server will create a new account if no matching
// identity exists; req.Username is used as the desired username in that
// case. req.Vars are forwarded to server runtime hooks as session vars and
// echoed back on the resulting session.
func (c *Client) AuthenticateCustom(ctx context.Context, customID string, req AuthenticateRequest) (*Session, error) {
	body, err := json.Marshal(struct {
		ID   string            `json:"id"`
		Vars map[string]string `json:"vars,omitempty"`
	}{ID: customID, Vars: req.Vars})
	if err != nil {
		return nil, fmt.Errorf("nkclient: marshal request: %w", err)
	}

	q := url.Values{}
	q.Set("create", fmt.Sprintf("%t", req.Create))
	if req.Username != "" {
		q.Set("username", req.Username)
	}
	endpoint := c.baseURL() + "/v2/account/authenticate/custom?" + q.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("nkclient: build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.SetBasicAuth(c.serverKey, "")

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, decodeAPIError(resp)
	}

	var payload struct {
		Created      bool   `json:"created"`
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("nkclient: decode response: %w", err)
	}

	return sessionFromTokens(payload.Created, payload.Token, payload.RefreshToken)
}

// baseURL returns the scheme://host:port prefix used for every REST call.
func (c *Client) baseURL() string {
	scheme := "http"
	if c.cfg.UseSSL {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s:%d", scheme, c.cfg.Host, c.cfg.Port)
}

func decodeAPIError(resp *http.Response) error {
	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&body)
	return &APIError{Status: resp.StatusCode, Code: body.Code, Message: body.Message}
}

type sessionClaims struct {
	UserID    string            `json:"uid"`
	Username  string            `json:"usn"`
	Vars      map[string]string `json:"vrs"`
	ExpiresAt int64             `json:"exp"`
}

func sessionFromTokens(created bool, token, refreshToken string) (*Session, error) {
	claims, err := decodeJWTClaims(token)
	if err != nil {
		return nil, fmt.Errorf("nkclient: decode token: %w", err)
	}

	sess := &Session{
		Token:        token,
		RefreshToken: refreshToken,
		UserID:       claims.UserID,
		Username:     claims.Username,
		Vars:         claims.Vars,
		Created:      created,
	}
	if claims.ExpiresAt != 0 {
		sess.ExpiresAt = time.Unix(claims.ExpiresAt, 0)
	}
	if refreshToken != "" {
		if rc, err := decodeJWTClaims(refreshToken); err == nil && rc.ExpiresAt != 0 {
			sess.RefreshAt = time.Unix(rc.ExpiresAt, 0)
		}
	}
	return sess, nil
}

// decodeJWTClaims base64-decodes the payload segment of a JWT and unmarshals
// the Nakama session claims. The signature is not verified; the server is
// trusted to have signed it.
func decodeJWTClaims(token string) (sessionClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return sessionClaims{}, fmt.Errorf("invalid token: expected 3 segments, got %d", len(parts))
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return sessionClaims{}, fmt.Errorf("decode payload: %w", err)
	}
	var c sessionClaims
	if err := json.Unmarshal(raw, &c); err != nil {
		return sessionClaims{}, fmt.Errorf("unmarshal claims: %w", err)
	}
	return c, nil
}
