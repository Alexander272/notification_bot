package message

import (
	"net/http"

	"github.com/Alexander272/si_bot/internal/models"
	"github.com/Alexander272/si_bot/internal/models/response"
	"github.com/Alexander272/si_bot/internal/services"
	"github.com/gin-gonic/gin"
)

type MessageHandlers struct {
	service services.Message
}

func NewMessageHandlers(service services.Message) *MessageHandlers {
	return &MessageHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Message) {
	handlers := NewMessageHandlers(service)

	messages := api.Group("send")
	{
		messages.POST("/notification", handlers.notification)
		messages.POST("/command", handlers.command)
	}
}

func (h *MessageHandlers) notification(c *gin.Context) {
	var dto models.Notification
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.SendList(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Сообщение отправлено"})
}

func (h *MessageHandlers) command(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.IdResponse{Message: "Not implemented"})
}
