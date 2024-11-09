package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
)

// define app dependency
type App struct {
	Router *gin.Engine
	config *config.Config
}

func New(config *config.Config) *App {
	app := &App{
		config: config,
	}
	app.SetupRoutes()
	return app
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.Port),
		Handler: app.Router,
	}
	log.Printf("Starting server on %s", app.config.Port)
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
		log.Println("server cancel")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
