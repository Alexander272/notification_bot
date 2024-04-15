package error_bot

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Alexander272/si_bot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type Message struct {
	Service Service     `json:"service" binding:"required"`
	Data    MessageData `json:"data" binding:"required"`
}

type Service struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type MessageData struct {
	Date    string `json:"date" binding:"required"`
	Error   string `json:"error" binding:"required"`
	IP      string `json:"ip" binding:"required"`
	URL     string `json:"url" binding:"required"`
	Request string `json:"request"`
}

func Send(c *gin.Context, e string, request interface{}) {
	var req []byte
	if request != nil {
		var err error
		req, err = json.Marshal(request)
		if err != nil {
			logger.Errorf("failed to marshal request body. error: %s", err.Error())
		}
	}

	message := Message{
		Service: Service{
			Id:   "bot",
			Name: "Notification Bot",
		},
		Data: MessageData{
			Date:    time.Now().Format("02/01/2006 - 15:04:05"),
			IP:      c.ClientIP(),
			URL:     fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String()),
			Error:   e,
			Request: string(req),
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(message); err != nil {
		logger.Errorf("failed to encode struct. error: %s", err.Error())
	}

	url := os.Getenv("ERR_URL")
	if url == "" {
		return
	}

	_, err := http.Post(url, "application/json", &buf)
	if err != nil {
		logger.Errorf("failed to send error to bot. error: %s", err.Error())
	}
}
