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

// App define app dependency.
type App struct {
	Router *gin.Engine
	cfg    *config.Config
}

// New - app constructor.
func New(ctx context.Context, cfg *config.Config) *App {
	app := &App{
		cfg: cfg,
	}
	app.SetupRoutes(ctx)
	return app
}

// Start - app 啟動.
func (app *App) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.cfg.Port),
		Handler:           app.Router,
		ReadHeaderTimeout: time.Minute,
	}
	log.Info(fmt.Sprintf("Starting server on %s", app.cfg.Port))
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
		log.Info("server cancel")
		timeout, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		err := server.Shutdown(timeout)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}
}
