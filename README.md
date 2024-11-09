# golang-rest-api-sample

This repository is for use golang to implemenation resful api

## setup with server

```golang
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
```

## setup application
```golang
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/application"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
)

func main() {
	app := application.New(config.AppConfig)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
```
## setup router
```golang
package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoute(router *gin.RouterGroup) {
	router.POST("", h.CreateNews)
	router.GET("", h.GetAllNews)
	router.GET("/:id", h.GetNewsByID)
	router.PUT("/:id", h.UpdateNewsByID)
	router.DELETE("/:id", h.DeleteNewsByID)
}

func (h *Handler) CreateNews(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (h *Handler) GetAllNews(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (h *Handler) GetNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (h *Handler) UpdateNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

func (h *Handler) DeleteNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
```
## setup test
```golang
package news_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/application"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
	"github.com/stretchr/testify/assert"
)

func Test_CreateNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange test
			app := application.New(config.AppConfig)
			app.SetupRoutes()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/news", nil)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equalf(t, tc.expectedStatus, w.Code, "expected %v ,got %v", tc.expectedStatus, w.Code)
		})
	}
}

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange test
			app := application.New(config.AppConfig)
			app.SetupRoutes()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/news", nil)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equalf(t, tc.expectedStatus, w.Code, "expected %v ,got %v", tc.expectedStatus, w.Code)
		})
	}
}

func Test_GetNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange test
			app := application.New(config.AppConfig)
			app.SetupRoutes()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/news/1", nil)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equalf(t, tc.expectedStatus, w.Code, "expected %v ,got %v", tc.expectedStatus, w.Code)
		})
	}
}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange test
			app := application.New(config.AppConfig)
			app.SetupRoutes()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPut, "/news/1", nil)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equalf(t, tc.expectedStatus, w.Code, "expected %v ,got %v", tc.expectedStatus, w.Code)
		})
	}
}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange test
			app := application.New(config.AppConfig)
			app.SetupRoutes()

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/news/1", nil)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equalf(t, tc.expectedStatus, w.Code, "expected %v ,got %v", tc.expectedStatus, w.Code)
		})
	}
}
```