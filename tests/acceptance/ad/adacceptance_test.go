package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"os"
	"testing"
	"time"
)

var (
	srv             server.Server
	dependencies, _ = configuration.BuildAdConfiguration()
)

func TestMain(m *testing.M) {
	_, srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, dependencies)
	exitCode := m.Run()
	os.Exit(exitCode)
}
