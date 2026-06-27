package integration

import (
	"context"
	"fmt"
	"testing"
	"tests/containers"
	"tests/dbq"
	"time"

	"github.com/dsolerh/nkclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAuthenticateCustom exercises the custom-id authentication flow end to
// end against a real Nakama server. Subtests share a single container; each
// case uses unique credentials so they stay independent.
func TestAuthenticateCustom(t *testing.T) {
	ctx := t.Context()

	nk := containers.StartNakamaContainer(t, ctx)
	client := newTestClient(t, nk)
	db := nk.OpenDB(t, ctx)

	t.Run("when create is true and custom id is new should create a fresh session", func(t *testing.T) {
		subCtx := t.Context()

		customID := uniqueID("new")
		username := uniqueID("user")
		session, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create:   true,
			Username: username,
		})

		require.NoError(t, err)
		require.NotNil(t, session)
		assert.NotEmpty(t, session.Token, "token must be populated on success")
		assert.NotEmpty(t, session.UserID, "user id must be populated on success")
		assert.True(t, session.Created, "Created must be true for a brand-new account")
		assert.False(t, session.Expired(), "fresh session must not be expired")

		user := dbq.UserByCustomID(t, subCtx, db, customID)
		require.NotNil(t, user, "users row must exist for the new custom id")
		assert.Equal(t, session.UserID, user.ID, "users.id must match the session user id")
		assert.Equal(t, username, user.Username, "users.username must match the requested username")
		assert.Equal(t, customID, user.CustomID.String, "users.custom_id must be set to the request id")
		assert.True(t, user.CustomID.Valid, "users.custom_id must not be NULL")
		assert.WithinDuration(t, time.Now(), user.CreateTime, time.Minute, "create_time should be ~now")
		assert.True(t, user.DisableTime.Year() <= 1970, "freshly created user must not be disabled")
	})

	t.Run("when re-authenticating an existing custom id should reuse the account", func(t *testing.T) {
		subCtx := t.Context()

		customID := uniqueID("existing")
		username := uniqueID("user")

		first, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create:   true,
			Username: username,
		})
		require.NoError(t, err)
		require.NotNil(t, first)
		require.True(t, first.Created)
		require.Equal(t, 1, dbq.UserCountByCustomID(t, subCtx, db, customID),
			"exactly one users row after initial create")

		second, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create: true,
		})
		require.NoError(t, err)
		require.NotNil(t, second)
		assert.False(t, second.Created, "Created must be false when reusing an account")
		assert.Equal(t, first.UserID, second.UserID, "same custom id must resolve to the same user id")

		assert.Equal(t, 1, dbq.UserCountByCustomID(t, subCtx, db, customID),
			"re-auth must not create a duplicate users row")
		user := dbq.UserByCustomID(t, subCtx, db, customID)
		require.NotNil(t, user)
		assert.Equal(t, first.UserID, user.ID, "users.id must still match the original user id")
		assert.Equal(t, username, user.Username, "re-auth must not overwrite the username")
	})

	t.Run("when create is false and account does not exist should return an error", func(t *testing.T) {
		subCtx := t.Context()

		customID := uniqueID("missing")
		session, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create: false,
		})

		require.Error(t, err)
		assert.Nil(t, session, "session must be nil when authentication fails")
		assert.Equal(t, 0, dbq.UserCountByCustomID(t, subCtx, db, customID),
			"failed auth must not create a users row")
	})

	t.Run("when custom id is too short should return an error", func(t *testing.T) {
		subCtx := t.Context()

		customID := "abc"
		session, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create: true,
		})

		require.Error(t, err)
		assert.Nil(t, session)
		assert.Equal(t, 0, dbq.UserCountByCustomID(t, subCtx, db, customID),
			"invalid custom id must not create a users row")
	})

	t.Run("when vars are supplied should propagate them onto the session", func(t *testing.T) {
		subCtx := t.Context()

		customID := uniqueID("vars")
		vars := map[string]string{"region": "eu", "tier": "gold"}
		session, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create:   true,
			Username: uniqueID("user"),
			Vars:     vars,
		})

		require.NoError(t, err)
		require.NotNil(t, session)
		assert.Equal(t, vars, session.Vars, "session vars must echo the request vars")

		user := dbq.UserByCustomID(t, subCtx, db, customID)
		require.NotNil(t, user, "users row must exist for the new custom id")
		assert.Equal(t, "{}", user.Metadata,
			"session vars must not be persisted onto users.metadata")
	})

	t.Run("when context is cancelled should return an error", func(t *testing.T) {
		subCtx, cancel := context.WithCancel(t.Context())
		cancel()

		customID := uniqueID("cancel")
		session, err := client.AuthenticateCustom(subCtx, customID, nkclient.AuthenticateRequest{
			Create: true,
		})

		require.ErrorIs(t, err, context.Canceled)
		assert.Nil(t, session)
		assert.Equal(t, 0, dbq.UserCountByCustomID(t, ctx, db, customID),
			"cancelled auth must not create a users row")
	})
}

func newTestClient(t *testing.T, nk containers.NakamaContainer) *nkclient.Client {
	t.Helper()
	client, err := nkclient.NewClient(
		"defaultkey",
		nk.Host,
		nkclient.WithPort(nk.APIPort),
		nkclient.WithTimeout(10*time.Second),
	)
	require.NoError(t, err)
	return client
}

// uniqueID produces a per-test identifier that is long enough to satisfy
// Nakama's minimum length requirements while keeping subtests isolated.
func uniqueID(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}
