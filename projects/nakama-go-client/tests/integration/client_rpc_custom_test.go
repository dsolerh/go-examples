package integration

import (
	"context"
	_ "embed"
	"testing"
	"tests/containers"

	"github.com/dsolerh/nkclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/echo_plugin.js
var echoPluginJS string

// TestCallCustomRPC exercises (*Client).CallCustomRPC against a Nakama
// server loaded with a JS runtime plugin that exposes a handful of RPCs
// covering the happy path, server-side failures, and unknown ids.
func TestCallCustomRPC(t *testing.T) {
	ctx := t.Context()

	nk := containers.StartNakamaContainer(t, ctx,
		containers.WithJSPlugin("echo_plugin.js", echoPluginJS),
	)
	client := newTestClient(t, nk)

	session, err := client.AuthenticateCustom(ctx, uniqueID("rpc"), nkclient.AuthenticateRequest{
		Create:   true,
		Username: uniqueID("user"),
	})
	require.NoError(t, err, "authenticate before invoking rpcs")
	require.NotNil(t, session)

	t.Run("when echo rpc is invoked should round-trip the payload", func(t *testing.T) {
		subCtx := t.Context()

		type echoPayload struct {
			Greeting string `json:"greeting"`
			Count    int    `json:"count"`
		}
		req := echoPayload{Greeting: "hello, nakama", Count: 3}

		resp, err := nkclient.CallCustomRPC[echoPayload, echoPayload](subCtx, client, session, "echo", req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, req, *resp, "echo must return the exact request payload")
	})

	t.Run("when uppercase rpc is invoked should return the transformed response", func(t *testing.T) {
		subCtx := t.Context()

		type uppercaseReq struct {
			Text string `json:"text"`
		}
		type uppercaseResp struct {
			Result string `json:"result"`
		}

		resp, err := nkclient.CallCustomRPC[uppercaseReq, uppercaseResp](subCtx, client, session, "uppercase", uppercaseReq{Text: "hello world"})
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, "HELLO WORLD", resp.Result)
	})

	t.Run("when rpc throws on the server should surface an APIError", func(t *testing.T) {
		subCtx := t.Context()

		resp, err := nkclient.CallCustomRPC[any, any](subCtx, client, session, "fail", nil)
		require.Error(t, err)
		assert.Nil(t, resp)

		var apiErr *nkclient.APIError
		assert.ErrorAs(t, err, &apiErr, "server failures must be wrapped in *APIError")
	})

	t.Run("when rpc id is unknown should surface an APIError", func(t *testing.T) {
		subCtx := t.Context()

		resp, err := nkclient.CallCustomRPC[any, any](subCtx, client, session, "does-not-exist", nil)
		require.Error(t, err)
		assert.Nil(t, resp)

		var apiErr *nkclient.APIError
		assert.ErrorAs(t, err, &apiErr)
	})

	t.Run("when context is cancelled should return an error", func(t *testing.T) {
		subCtx, cancel := context.WithCancel(t.Context())
		cancel()

		err := client.CallCustomRPC(subCtx, session, "echo", map[string]string{"x": "y"}, nil)
		require.ErrorIs(t, err, context.Canceled)
	})
}
