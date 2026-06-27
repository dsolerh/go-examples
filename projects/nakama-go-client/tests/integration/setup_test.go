package integration

import (
	"context"
	"os"
	"testing"
	"tests/containers"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	containers.SetupPgInstance(ctx)

	code := m.Run()

	containers.ClosePgInstance(ctx)
	os.Exit(code)
}
