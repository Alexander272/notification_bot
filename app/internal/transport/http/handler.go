package http

import (
	"github.com/Alexander272/si_bot/internal/config"
	"github.com/Alexander272/si_bot/pkg/limiter"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		limiter.Limit(conf.Limiter.RPS, conf.Limiter.Burst, conf.Limiter.TTL),
	)

	// TODO init api

	return router
}
