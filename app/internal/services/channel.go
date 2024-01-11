package services

import (
	"context"
	"fmt"

	"github.com/mattermost/mattermost-server/v6/model"
)

type ChannelService struct {
	MostClient *model.Client4
	BotName    string
}

func NewChannelService(client *model.Client4, botName string) *ChannelService {
	return &ChannelService{
		MostClient: client,
		BotName:    botName,
	}
}

type Channel interface {
	Create(ctx context.Context, userId string) (string, error)
}

// create new direct chanel
func (s *ChannelService) Create(ctx context.Context, userId string) (string, error) {
	//TODO возможно лучше делать это в другом месте тк id бота понадобится в других местах
	bot, _, err := s.MostClient.GetUserByUsername(s.BotName, "")
	if err != nil {
		return "", fmt.Errorf("failed to get bot. error: %w", err)
	}

	channel, _, err := s.MostClient.CreateDirectChannel(bot.Id, userId)
	if err != nil {
		return "", fmt.Errorf("failed to create direct channel. error: %w", err)
	}

	return channel.Id, nil
}
