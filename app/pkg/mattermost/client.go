package mattermost

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerLink string
	Token      string
}

type Client struct {
	Http   *model.Client4
	Socket *model.WebSocketClient
}

// может разделить клиенты и отдельно подключать http и websocket
func NewMattermostClient(conf Config) *Client {
	httpClient := model.NewAPIv4Client("http://" + conf.ServerLink)
	httpClient.SetToken(conf.Token)

	socketClient, err := model.NewWebSocketClient4("ws://"+conf.ServerLink, conf.Token)
	if err != nil {
		logrus.Fatalf("failed to websocket connect to mattermost. error: %s", err.Error())
	}

	return &Client{
		Http:   httpClient,
		Socket: socketClient,
	}
}
