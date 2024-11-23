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

## Add Mage Tools

[Mage](https://magefile.org/) is a tool for setup build tool with golang

### how to use

1. install dependency

```shell
go get github.com/magefile/mage/mage
```

2. create a execute file

```golang
//go:build ignore
// +build ignore

package main

import (
	"os"

	"github.com/magefile/mage/mage"
)

func main() { os.Exit(mage.Main()) }
```

**Notice**: 檔案前面的 // +build ignore 是避免 mage 編譯這隻檔案

3. create a task file for run task

```golang
//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// clean the build binary
func Clean() error {
	return sh.Rm("bin")
}

// Creates the binary in the current directory.
func Build() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/server", "./cmd/main.go")
}

// start the server
func Launch() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/server")
	if err != nil {
		return err
	}
	return nil
}

// run the test
func Test() error {
	err := sh.RunV("go", "test", "-v", "./...")
	if err != nil {
		return err
	}
	return nil
}
```

task 形式分成兩種 1. 有 error return 值 2. 沒有 error return 

而參數的部份可以帶入 context 來作而外的處理

相依性則透過 mg.Deps 來設定

預設不會列印執行結果及過程。如果需要有列印執行結果，需要使用 sh.RunV 來作處理