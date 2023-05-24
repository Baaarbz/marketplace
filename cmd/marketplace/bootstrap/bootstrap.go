package bootstrap

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"context"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type config struct {
	// API configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8000"`
	ShutdownTimeout time.Duration `default:"10s"`
}

var cfg config

func Run() error {
	// Load config
	if err := envconfig.Process("marketplace", &cfg); err != nil {
		return err
	}

	adConfiguration, err := configuration.BuildAdConfiguration()
	if err != nil {
		return err
	}

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, adConfiguration)
	return srv.Run(ctx)
}
