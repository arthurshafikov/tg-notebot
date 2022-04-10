package handler

import (
	"context"
	"net/http"

	"github.com/arthurshafikov/tg-notebot/internal/core"
	"github.com/arthurshafikov/tg-notebot/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctx      context.Context
	services *services.Services
}

func NewHandler(
	ctx context.Context,
	services *services.Services,
) *Handler {
	return &Handler{
		ctx:      ctx,
		services: services,
	}
}

func (h *Handler) Init(e *gin.Engine) {
}

func (h *Handler) setUnprocessableEntityJSONResponse(ctx *gin.Context, data string) {
	h.setJSONResponse(ctx, http.StatusUnprocessableEntity, data)
}

func (h *Handler) setJSONResponse(ctx *gin.Context, code int, data string) {
	ctx.JSON(code, core.ServerResponse{
		Data: data,
	})
}
