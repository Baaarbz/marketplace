package bootstrap

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"time"
)

type config struct {
	// API configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8000"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"barbz" required:"true"`
	DbPass    string        `default:"1234" required:"true"`
	DbHost    string        `default:"localhost" required:"true"`
	DbPort    string        `default:"5432" required:"true"`
	DbName    string        `default:"marketplace" required:"true"`
	DbTimeout time.Duration `default:"5s"`
}

var cfg config

func Run() error {
	// Load config
	if err := envconfig.Process("marketplace", &cfg); err != nil {
		return err
	}

	db, err := initDatabase()
	if err != nil {
		return err
	}

	err = initGoose(db)
	if err != nil {
		return err
	}

	adConfiguration, err := configuration.BuildAdConfiguration(db, cfg.ShutdownTimeout)
	if err != nil {
		return err
	}

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, adConfiguration)
	return srv.Run(ctx)
}

func initDatabase() (*sql.DB, error) {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName)

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

func initGoose(db *sql.DB) error {
	err := goose.Up(db, "db/migrations")
	return err
}
