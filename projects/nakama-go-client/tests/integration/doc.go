// Package integration runs nkclient against a real Nakama server spun up
// via testcontainers-go.
//
// All tests in this package are guarded by the `integration` build tag
// because they require a running Docker daemon and pull/spawn Postgres
// and Nakama containers. Run with:
//
//	go test -tags=integration ./integration/...
//
// A single Postgres container is shared across the whole suite. Each
// Nakama container connects to a freshly-created database inside that
// Postgres instance so tests stay isolated from each other.
package integration
