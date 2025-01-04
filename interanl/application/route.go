package application

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/logger"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/services/news"
)

// SetupRoutes - define route.
func (app *App) SetupRoutes(ctx context.Context) {
	gin.SetMode(app.cfg.GinMode)
	router := gin.New()
	// recovery middleward
	router.Use(sloggin.New(logger.FromContext(ctx)))
	router.Use(gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
	})
	app.Router = router
	app.SetupNewsRoutes()
}

// SetupNewsRoutes - setup new routes.
func (app *App) SetupNewsRoutes() {
	newGroups := app.Router.Group("/news")
	newsStore := news.NewNewsStore()
	handler := news.NewHandler(newsStore)
	handler.RegisterRoute(newGroups)
}
