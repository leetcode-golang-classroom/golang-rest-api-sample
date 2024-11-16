package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/logger"
)

// define app dependency
type App struct {
	Router *gin.Engine
	config *config.Config
}

func New(config *config.Config, ctx context.Context) *App {
	app := &App{
		config: config,
	}
	app.SetupRoutes(ctx)
	return app
}

func (app *App) Start(ctx context.Context) error {
	logger := logger.FromContext(ctx)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.Port),
		Handler: app.Router,
	}
	logger.Info(fmt.Sprintf("Starting server on %s", app.config.Port))
	errCh := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		logger.Info("server cancel")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
