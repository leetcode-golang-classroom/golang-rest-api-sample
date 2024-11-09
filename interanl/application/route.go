package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/services/news"
)

// define route
func (app *App) SetupRoutes() {
	gin.SetMode(app.config.GinMode)
	router := gin.New()
	// recovery middleward
	router.Use(gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "ok"})
	})
	app.Router = router
	app.SetupNewsRoutes()
}

func (app *App) SetupNewsRoutes() {
	newGroups := app.Router.Group("/news")
	handler := news.NewHandler()
	handler.RegisterRoute(newGroups)
}
