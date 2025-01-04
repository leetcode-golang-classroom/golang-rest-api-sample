package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler - Handler type.
type Handler struct {
	ns Storer
}

// NewHandler - create a Handler.
func NewHandler(ns Storer) *Handler {
	return &Handler{
		ns: ns,
	}
}

// RegisterRoute - register route.
func (h *Handler) RegisterRoute(router *gin.RouterGroup) {
	router.POST("", h.CreateNews)
	router.GET("", h.GetAllNews)
	router.GET("/:id", h.GetNewsByID)
	router.PUT("/:id", h.UpdateNewsByID)
	router.DELETE("/:id", h.DeleteNewsByID)
}

// CreateNews - create news handler.
func (h *Handler) CreateNews(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

// GetAllNews - get all news route.
func (h *Handler) GetAllNews(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

// GetNewsByID - get specific news route.
func (h *Handler) GetNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

// UpdateNewsByID - update news route.
func (h *Handler) UpdateNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}

// DeleteNewsByID - delete news route.
func (h *Handler) DeleteNewsByID(ctx *gin.Context) {
	ctx.Status(http.StatusNotImplemented)
}
