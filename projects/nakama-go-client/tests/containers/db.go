package containers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"sync/atomic"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tcpg "github.com/testcontainers/testcontainers-go/modules/postgres"
	tcnet "github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	pgImage    = "postgres:16-alpine"
	pgUser     = "postgres"
	pgPassword = "password"
	pgAlias    = "postgres"
)

var (
	sharedNet *testcontainers.DockerNetwork
	pgCont    *tcpg.PostgresContainer
	dbCounter atomic.Uint64
)

func SetupPgInstance(ctx context.Context) {

	net, err := tcnet.New(ctx)
	if err != nil {
		log.Fatalf("create docker network: %v", err)
	}
	sharedNet = net

	pg, err := tcpg.Run(ctx,
		pgImage,
		tcpg.WithUsername(pgUser),
		tcpg.WithPassword(pgPassword),
		tcpg.WithDatabase("postgres"),
		tcnet.WithNetwork([]string{pgAlias}, sharedNet),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second),
		),
	)
	if err != nil {
		_ = sharedNet.Remove(ctx)
		log.Fatalf("start postgres: %v", err)
	}
	pgCont = pg
}

func ClosePgInstance(ctx context.Context) {
	_ = pgCont.Terminate(ctx)
	_ = sharedNet.Remove(ctx)
}

// uniqueDB provisions a fresh, empty database inside the shared Postgres
// instance and returns its name. Each call yields a different database so
// individual Nakama containers stay isolated from one another.
func uniqueDB(t *testing.T, ctx context.Context) string {
	t.Helper()

	name := fmt.Sprintf("nakama_test_%d", dbCounter.Add(1))
	exitCode, out, err := pgCont.Exec(ctx, []string{
		"psql", "-U", pgUser, "-d", "postgres",
		"-c", fmt.Sprintf("CREATE DATABASE %s", name),
	})
	require.NoErrorf(t, err, "create db %s", name)
	if exitCode != 0 {
		body, _ := io.ReadAll(out)
		t.Fatalf("create db %s: exit %d: %s", name, exitCode, string(body))
	}
	return name
}

// hostDSN returns a libpq-style DSN that reaches the shared Postgres
// container from the test process (i.e. via the host-mapped port, not the
// docker network alias the Nakama container uses).
func hostDSN(t *testing.T, ctx context.Context, dbName string) string {
	t.Helper()

	host, err := pgCont.Host(ctx)
	require.NoError(t, err, "postgres host")
	port, err := pgCont.MappedPort(ctx, "5432/tcp")
	require.NoError(t, err, "postgres mapped port")
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser, pgPassword, host, port.Port(), dbName,
	)
}

// OpenDB opens a *sql.DB connected to the named database inside the shared
// Postgres instance from the host network. The connection is closed when
// the test ends.
func OpenDB(t *testing.T, ctx context.Context, dbName string) *sql.DB {
	t.Helper()

	db, err := sql.Open("pgx", hostDSN(t, ctx, dbName))
	require.NoErrorf(t, err, "open db %s", dbName)
	t.Cleanup(func() { _ = db.Close() })
	require.NoErrorf(t, db.PingContext(ctx), "ping db %s", dbName)
	return db
}
