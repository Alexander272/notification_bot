package services

import "github.com/mattermost/mattermost-server/v6/model"

type Services struct {
	Channel
	Post
	Dialog
}

type Deps struct {
	MostClient *model.Client4
	BotName    string
}

func NewServices(deps Deps) *Services {
	channel := NewChannelService(deps.MostClient, deps.BotName)
	post := NewPostService(deps.MostClient, channel)
	dialog := NewDialogService(deps.MostClient)

	return &Services{
		Channel: channel,
		Post:    post,
		Dialog:  dialog,
	}
}
