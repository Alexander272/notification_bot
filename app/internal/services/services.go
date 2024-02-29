package services

import "github.com/mattermost/mattermost-server/v6/model"

type Services struct {
	Channel
	Message
}

type Deps struct {
	MostClient *model.Client4
	BotName    string
	AppUrl     string
}

func NewServices(deps Deps) *Services {
	channel := NewChannelService(deps.MostClient, deps.BotName)
	message := NewMessageService(deps.MostClient, deps.AppUrl, channel)

	return &Services{
		Channel: channel,
		Message: message,
	}
}
