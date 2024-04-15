package http

import (
	"github.com/Alexander272/si_bot/internal/config"
	"github.com/Alexander272/si_bot/internal/services"
	"github.com/Alexander272/si_bot/internal/transport/http/dialogs"
	"github.com/Alexander272/si_bot/internal/transport/http/posts"
	"github.com/Alexander272/si_bot/pkg/limiter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		limiter.Limit(conf.Limiter.RPS, conf.Limiter.Burst, conf.Limiter.TTL),
	)

	api := router.Group("/api")

	posts.Register(api, h.services.Post)
	dialogs.Register(api, h.services.Dialog)

	return router
}
