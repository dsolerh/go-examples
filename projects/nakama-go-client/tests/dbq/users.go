// Package dbq holds query helpers that integration tests use to assert on
// Nakama's Postgres state. Each helper takes a *sql.DB so callers stay in
// control of which database (i.e. which Nakama container) is being checked.
package dbq

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// User is a partial projection of Nakama's `users` table, scoped to the
// columns integration tests actually assert on.
type User struct {
	ID          string
	Username    string
	CustomID    sql.NullString
	Metadata    string
	CreateTime  time.Time
	UpdateTime  time.Time
	DisableTime time.Time
}

const userSelect = `SELECT id::text, username, custom_id, metadata::text,
	create_time, update_time, disable_time
FROM users`

// UserByCustomID returns the user with the given custom_id, or nil if no
// such row exists. Any other query error fails the test.
func UserByCustomID(t *testing.T, ctx context.Context, db *sql.DB, customID string) *User {
	t.Helper()

	var u User
	err := db.QueryRowContext(ctx, userSelect+" WHERE custom_id = $1", customID).
		Scan(&u.ID, &u.Username, &u.CustomID, &u.Metadata,
			&u.CreateTime, &u.UpdateTime, &u.DisableTime)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	require.NoError(t, err, "query user by custom_id")
	return &u
}

// UserCountByCustomID returns the number of users matching custom_id.
func UserCountByCustomID(t *testing.T, ctx context.Context, db *sql.DB, customID string) int {
	t.Helper()

	var n int
	require.NoError(t,
		db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE custom_id = $1`, customID).Scan(&n),
		"count users by custom_id",
	)
	return n
}
