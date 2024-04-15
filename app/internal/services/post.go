package services

import (
	"context"
	"fmt"

	"github.com/Alexander272/si_bot/internal/models"
	"github.com/mattermost/mattermost-server/v6/model"
)

type PostService struct {
	MostClient *model.Client4
	channel    Channel
}

func NewPostService(mostClient *model.Client4, channel Channel) *PostService {
	return &PostService{
		MostClient: mostClient,
		channel:    channel,
	}
}

type Post interface {
	SendPost(context.Context, models.CreatePostDTO) error
	UpdatePost(context.Context, models.UpdatePostDTO) error
}

func (s *PostService) SendPost(ctx context.Context, data models.CreatePostDTO) error {
	post := &model.Post{
		ChannelId: data.ChannelId,
		Message:   data.Message,
	}

	if data.UserId != "" {
		channelId, err := s.channel.Create(ctx, data.UserId)
		if err != nil {
			return err
		}
		post.ChannelId = channelId
	}

	if data.Actions != nil {
		attachment := &model.SlackAttachment{
			Actions: data.Actions,
		}
		post.AddProp("attachments", []*model.SlackAttachment{attachment})
	}
	if data.Attachments != nil {
		post.AddProp("attachments", data.Attachments)
	}

	for _, p := range data.Props {
		post.AddProp(p.Key, p.Value)
	}

	//TODO
	// можно передавать ID поста. выполнять поиск в канале этого ID, если он есть удалять сообщение и отправлять новое
	// благодаря такой схеме можно избежать варианта, когда в канале три одинаковых сообщений (для получения инструмента) с разной датой, а изменяться по нажатию будет только последнее
	// если я задаю ID метод выдает ошибку (только какого хрена он позволяет его задавать)
	// если записывать ID поста в props и закреплять сообщение, а потом получать все закрепленные и искать в props нужный мне ID
	// в ID поста надо учитывать ID пользователя из-за которого этот пост создавался или что-то подобное (у метролога может быть много всего)

	_, _, err := s.MostClient.CreatePost(post)
	if err != nil {
		return fmt.Errorf("failed to create post. error: %w", err)
	}
	return nil
}

func (s *PostService) UpdatePost(ctx context.Context, data models.UpdatePostDTO) error {
	// тут тогда надо будет откреплять пост (скорее всего, можно будет передавать флаг)
	post := &model.Post{
		Id:       data.PostId,
		Message:  data.Message,
		IsPinned: false,
	}

	if data.Actions != nil {
		attachment := &model.SlackAttachment{
			Actions: data.Actions,
		}
		post.AddProp("attachments", []*model.SlackAttachment{attachment})
	}
	if data.Attachments != nil {
		post.AddProp("attachments", data.Attachments)
	}

	for _, p := range data.Props {
		post.AddProp(p.Key, p.Value)
	}

	_, _, err := s.MostClient.UpdatePost(data.PostId, post)
	if err != nil {
		return fmt.Errorf("failed to create post. error: %w", err)
	}
	return nil
}
