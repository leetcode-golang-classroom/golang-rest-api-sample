package news_test

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/application"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/logger"
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
			ctx := context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
			})))
			// Arrange test
			app := application.New(ctx, config.AppConfig)
			app.SetupRoutes(ctx)

			w := httptest.NewRecorder()
			r, err := http.NewRequestWithContext(ctx, http.MethodPost, "/news", http.NoBody)
			require.NoError(t, err)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equal(t, tc.expectedStatus, w.Code)
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
			ctx := context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
			})))
			app := application.New(ctx, config.AppConfig)
			app.SetupRoutes(ctx)

			w := httptest.NewRecorder()
			r, err := http.NewRequestWithContext(ctx, http.MethodGet, "/news", http.NoBody)
			require.NoError(t, err)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equal(t, tc.expectedStatus, w.Code)
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
			ctx := context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
			})))
			app := application.New(ctx, config.AppConfig)
			app.SetupRoutes(ctx)

			w := httptest.NewRecorder()
			r, err := http.NewRequestWithContext(ctx, http.MethodGet, "/news/1", http.NoBody)
			require.NoError(t, err)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equal(t, tc.expectedStatus, w.Code)
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
			ctx := context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
			})))
			app := application.New(ctx, config.AppConfig)
			app.SetupRoutes(ctx)

			w := httptest.NewRecorder()
			r, err := http.NewRequestWithContext(ctx, http.MethodPut, "/news/1", http.NoBody)
			require.NoError(t, err)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equal(t, tc.expectedStatus, w.Code)
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
			ctx := context.WithValue(context.Background(), logger.CtxKey{}, slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
			})))
			app := application.New(ctx, config.AppConfig)
			app.SetupRoutes(ctx)

			w := httptest.NewRecorder()
			r, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/news/1", http.NoBody)
			require.NoError(t, err)
			// Act
			app.Router.ServeHTTP(w, r)
			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
