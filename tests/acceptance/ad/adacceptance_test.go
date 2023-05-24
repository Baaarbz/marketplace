package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

var (
	srv             server.Server
	db, _           = initDatabase()
	dependencies, _ = configuration.BuildAdConfiguration(db, 10*time.Second)
)

func TestMain(m *testing.M) {
	initDocker()
	_, srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, dependencies)
	exitCode := m.Run()
	stopDocker()
	os.Exit(exitCode)
}
func initDatabase() (*sql.DB, error) {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		"barbz",
		"1234",
		"localhost",
		"5432",
		"marketplace")

	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	// Ping DB to check if the connection was established successfully
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func initDocker() {
	cmd := exec.Command("docker-compose", "-p", "marketplace-acceptance-test", "up")

	// Channel to receive interrupt signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Exec docker compose in a gorutine
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func stopDocker() {
	cmd := exec.Command("docker-compose", "down", "-d")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
