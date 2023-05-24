package server

import (
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	"barbz.dev/marketplace/internal/infrastructure/server/handler/health"
	"barbz.dev/marketplace/internal/infrastructure/server/middleware/logging"
	"barbz.dev/marketplace/internal/infrastructure/server/middleware/recovery"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Engine          *gin.Engine
	HttpAddr        string
	ShutdownTimeout time.Duration
	// Dependencies
	AdConfiguration *configuration.AdConfiguration
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, adConfiguration *configuration.AdConfiguration) (context.Context, Server) {
	engine := gin.New()
	// Register middlewares
	engine.Use(recovery.Middleware(), logging.Middleware())

	srv := Server{
		Engine:          engine,
		HttpAddr:        fmt.Sprintf("%s:%d", host, port),
		ShutdownTimeout: shutdownTimeout,
		AdConfiguration: adConfiguration,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("server running on", s.HttpAddr)

	srv := &http.Server{
		Addr:    s.HttpAddr,
		Handler: s.Engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.Engine.GET("/health", health.APIStatus())

	saveAdHandler := ad.NewSaveAdHandler(s.AdConfiguration)
	getAdHandler := ad.NewGetAdHandler(s.AdConfiguration)
	adsGroup := s.Engine.Group("api/v1/ads")
	{
		adsGroup.POST("", saveAdHandler.SaveAd())
		adsGroup.GET("", getAdHandler.FindAllAds())
		adsGroup.GET(":id", getAdHandler.FindAdById())
	}
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
