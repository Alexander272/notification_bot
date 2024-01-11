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
	userId := "8h9j8ogsupd83f3k6jtr1ynhiw"

	list := []models.SI{
		{
			Id:            "026f772a-72f9-47a1-ac89-c73d65832c47",
			Name:          "Штангенциркуль цифровой",
			FactoryNumber: "HIJA027140597",
			Department:    "Отдел технического сервиса",
			Person:        "test user",
		},
		{
			Id:            "0e64f324-5644-4efe-be97-59b2ca2f9699",
			Name:          "Набор щупов",
			FactoryNumber: "600",
			Department:    "Отдел технического сервиса",
			Person:        "test user2",
		},
		{
			Id:            "238b1d7f-309b-4123-ba9b-05ce98abbe47",
			Name:          "Прибор комбинированный",
			FactoryNumber: "83490106",
			Department:    "Отдел технического сервиса",
			Person:        "test user3",
		},
	}

	dto := models.Notification{UserId: userId, Message: "Отправлены инструменты", SI: list}

	if err := h.service.SendList(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Сообщение отправлено"})
}

func (h *MessageHandlers) command(c *gin.Context) {}
