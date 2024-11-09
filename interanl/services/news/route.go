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
