package nkclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// CallCustomRPC is a generic helper around (*Client).CallCustomRPC that
// allocates a fresh Response value and returns it on success.
func CallCustomRPC[Request, Response any](ctx context.Context, client *Client, session *Session, rpc string, req Request) (*Response, error) {
	resp := new(Response)
	if err := client.CallCustomRPC(ctx, session, rpc, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CallCustomRPC invokes a server-registered RPC by id and decodes the
// returned payload into response.
//
// request is marshalled to JSON and forwarded as the RPC payload; pass nil
// to send an empty payload. response may be nil when the caller does not
// care about the returned body.
func (c *Client) CallCustomRPC(ctx context.Context, session *Session, rpc string, request, response any) error {
	if session == nil {
		return fmt.Errorf("nkclient: session is required")
	}

	payload := ""
	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("nkclient: marshal request: %w", err)
		}
		payload = string(b)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("nkclient: marshal body: %w", err)
	}

	endpoint := c.baseURL() + "/v2/rpc/" + url.PathEscape(rpc)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("nkclient: build request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+session.Token)

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return decodeAPIError(resp)
	}

	if response == nil {
		return nil
	}

	var envelope struct {
		ID      string `json:"id"`
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return fmt.Errorf("nkclient: decode response: %w", err)
	}
	if envelope.Payload == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(envelope.Payload), response); err != nil {
		return fmt.Errorf("nkclient: decode payload: %w", err)
	}
	return nil
}
