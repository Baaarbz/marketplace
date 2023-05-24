package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	srv             server.Server
	db, _           = initDatabase()
	dependencies, _ = configuration.BuildAdConfiguration(db, 10*time.Second)
)

func TestMain(m *testing.M) {
	_, srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, dependencies)
	exitCode := m.Run()
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
