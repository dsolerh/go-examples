package containers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const image = "registry.heroiclabs.com/heroiclabs/nakama:3.39.0"

type NakamaContainer struct {
	Host    string
	APIPort int
	// DBName is the Postgres database this Nakama instance writes to.
	// Tests use it via OpenDB to assert on persisted state.
	DBName string
}

// NakamaOption mutates the configuration applied to a Nakama container
// before it is started.
type NakamaOption func(*nakamaConfig)

type jsPlugin struct {
	filename string
	content  string
}

type nakamaConfig struct {
	jsPlugin *jsPlugin
}

// WithJSPlugin mounts a JavaScript runtime module into the Nakama container
// and registers it as the JS runtime entrypoint. content is the literal
// source of the plugin; filename is the basename it will be written as under
// the runtime modules directory.
func WithJSPlugin(filename, content string) NakamaOption {
	return func(c *nakamaConfig) {
		c.jsPlugin = &jsPlugin{filename: filename, content: content}
	}
}

func StartNakamaContainer(t *testing.T, ctx context.Context, opts ...NakamaOption) NakamaContainer {
	t.Helper()

	var cfg nakamaConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	dbName := uniqueDB(t, ctx)
	dbAddr := fmt.Sprintf("%s:%s@%s:5432/%s", pgUser, pgPassword, pgAlias, dbName)

	flags := []string{
		"--name nakama-test",
		"--database.address " + dbAddr,
		"--logger.level INFO",
		"--session.token_expiry_sec 7200",
	}

	var files []testcontainers.ContainerFile
	if cfg.jsPlugin != nil {
		files = append(files, testcontainers.ContainerFile{
			Reader:            strings.NewReader(cfg.jsPlugin.content),
			ContainerFilePath: "/nakama/data/modules/" + cfg.jsPlugin.filename,
			FileMode:          0o644,
		})
		flags = append(flags, "--runtime.js_entrypoint "+cfg.jsPlugin.filename)
	}

	cmd := fmt.Sprintf(
		"/nakama/nakama migrate up --database.address %s && exec /nakama/nakama %s",
		dbAddr, strings.Join(flags, " "),
	)

	req := testcontainers.ContainerRequest{
		Image:        image,
		Entrypoint:   []string{"/bin/sh", "-ecx"},
		Cmd:          []string{cmd},
		ExposedPorts: []string{"7349/tcp", "7350/tcp", "7351/tcp"},
		Networks:     []string{sharedNet.Name},
		Files:        files,
		// "Startup done" is logged and the ports begin listening before
		// Nakama can actually serve: the HTTP gateway (7350) accepts
		// connections while its own gRPC backend (127.0.0.1:7349) is still
		// coming up, so early requests return 503/code=14. /healthcheck is
		// proxied by the gateway through to the gRPC Healthcheck RPC, so a
		// 200 confirms the full gateway->gRPC path is ready.
		WaitingFor: wait.ForAll(
			wait.ForLog("Startup done"),
			wait.ForHTTP("/healthcheck").
				WithPort("7350/tcp").
				WithStatusCodeMatcher(func(status int) bool { return status == 200 }),
		).WithStartupTimeoutDefault(2 * time.Minute),
	}

	nk, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err, "start nakama")
	t.Cleanup(func() {
		// Dump the server logs when the test failed so the gRPC/runtime
		// side of a failure is visible next to the assertion error. Runs
		// before Terminate so the container is still around to read from.
		if t.Failed() {
			dumpNakamaLogs(t, nk)
		}
		_ = nk.Terminate(context.Background())
	})

	host, err := nk.Host(ctx)
	require.NoError(t, err, "nakama host")
	port, err := nk.MappedPort(ctx, "7350")
	require.NoError(t, err, "nakama mapped port")

	return NakamaContainer{
		Host:    host,
		APIPort: int(port.Num()),
		DBName:  dbName,
	}
}

// dumpNakamaLogs writes the container's stdout/stderr to a file under
// ./testlogs and reports the path via the test log. It is best-effort: any
// error is logged but not fatal, since it runs while a test has already
// failed.
func dumpNakamaLogs(t *testing.T, c testcontainers.Container) {
	t.Helper()

	rc, err := c.Logs(context.Background())
	if err != nil {
		t.Logf("nakama: could not read container logs: %v", err)
		return
	}
	defer rc.Close()

	dir := "testlogs"
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Logf("nakama: could not create %s: %v", dir, err)
		return
	}
	// Sanitize the test name (subtests contain '/') into a flat filename.
	name := strings.ReplaceAll(t.Name(), "/", "_")
	path := filepath.Join(dir, "nakama-"+name+".log")

	f, err := os.Create(path)
	if err != nil {
		t.Logf("nakama: could not create log file %s: %v", path, err)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, rc); err != nil {
		t.Logf("nakama: could not write log file %s: %v", path, err)
		return
	}
	t.Logf("nakama container logs written to %s", path)
}

// OpenDB returns a *sql.DB connected to the Postgres database this Nakama
// container writes to, for assertions on persisted state. The connection is
// closed when the test ends.
func (n NakamaContainer) OpenDB(t *testing.T, ctx context.Context) *sql.DB {
	t.Helper()
	return OpenDB(t, ctx, n.DBName)
}
