package posts

import (
	"net/http"

	"github.com/Alexander272/si_bot/internal/models"
	"github.com/Alexander272/si_bot/internal/models/response"
	"github.com/Alexander272/si_bot/internal/services"
	"github.com/Alexander272/si_bot/pkg/error_bot"
	"github.com/Alexander272/si_bot/pkg/logger"
	"github.com/gin-gonic/gin"
)

type PostHandlers struct {
	service services.Post
}

func NewPostHandlers(service services.Post) *PostHandlers {
	return &PostHandlers{
		service: service,
	}
}

func Register(api *gin.RouterGroup, service services.Post) {
	handlers := NewPostHandlers(service)

	posts := api.Group("posts")
	{
		posts.POST("", handlers.create)
		posts.PUT("/:id", handlers.update)
	}
}

func (h *PostHandlers) create(c *gin.Context) {
	var dto models.CreatePostDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}
	logger.Info("userId", dto.UserId)

	//TODO может стоит попытаться как-нибудь определять какой пользователь вызвал тот или иной запрос
	if err := h.service.SendPost(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Сообщение отправлено"})
}

func (h *PostHandlers) update(c *gin.Context) {
	var dto models.UpdatePostDTO
	if err := c.BindJSON(&dto); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error(), "Отправлены некорректные данные")
		return
	}

	if err := h.service.UpdatePost(c, dto); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла ошибка. "+err.Error())
		error_bot.Send(c, err.Error(), dto)
		return
	}

	c.JSON(http.StatusOK, response.IdResponse{Message: "Сообщение обновлено"})
}
