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
