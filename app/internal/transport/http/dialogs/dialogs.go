package dialogs

import (
	"net/http"

	"github.com/Alexander272/si_bot/internal/models"
	"github.com/Alexander272/si_bot/internal/models/response"
	"github.com/Alexander272/si_bot/internal/services"
	"github.com/Alexander272/si_bot/pkg/error_bot"
	"github.com/gin-gonic/gin"
)

type DialogsHandler struct {
	service services.Dialog
}

func NewDialogsHandler(service services.Dialog) *DialogsHandler {
	return &DialogsHandler{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Dialog) {
	handlers := NewDialogsHandler(service)

	dialogs := api.Group("dialogs")
	{
		dialogs.POST("", handlers.create)
	}
}

func (h *DialogsHandler) create(c *gin.Context) {
	var dto models.PostAction
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.Open(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Сообщение отправлено"})
}
