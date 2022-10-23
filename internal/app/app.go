package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"service-echo/config"
	v1 "service-echo/internal/entrypoint/http/v1"
	"service-echo/internal/usecase"
	"service-echo/internal/usecase/rewriter"
	"service-echo/pkg/httpserver"
	"service-echo/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Use Case
	echoUseCase := usecase.New(rewriter.New(cfg.Rewriter))

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, echoUseCase)

	log.Printf("swagger docs on  http://localhost:%v/swagger/index.html", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	l.Debug("Server shutdown")

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
