package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/si_bot/constants"
	"github.com/Alexander272/si_bot/internal/models"
	"github.com/Alexander272/si_bot/pkg/logger"
	"github.com/mattermost/mattermost-server/v6/model"
)

type MessageService struct {
	MostClient *model.Client4
	AppUrl     string
	channel    Channel
}

func NewMessageService(mostClient *model.Client4, appUrl string, channel Channel) *MessageService {
	return &MessageService{
		MostClient: mostClient,
		AppUrl:     appUrl,
		channel:    channel,
	}
}

type Message interface {
	SendList(ctx context.Context, notification models.Notification) error
}

//? можно этот сервис сделать основным (искать тут канал и отправлять сообщения. либо прямо тут описывать структуру, либо сделать отдельный сервис с шаблонами)

/*
	можно забирать список отправленных инструментов таким запросом

	SELECT i.id, name, factory_number, m.status, person, department, date_of_receiving, date_of_issue
		FROM public.si_movement_history AS m
		INNER JOIN instruments AS i ON instrument_id=i.id
		WHERE department='Отдел технического сервиса' AND m.status='moved' AND date_of_receiving=0
*/

// нужно отправлять списки когда пользователю отправляют инструменты, когда их надо сдать на поверку и еще можно отправлять по запросу
// send list si |
func (s *MessageService) SendList(ctx context.Context, notification models.Notification) error {
	table := []string{
		"| Наименование СИ | зав.№ | Держатель |",
		"|:--|:--|:--|",
	}

	instrumentIds := []string{}

	for _, si := range notification.SI {
		instrumentIds = append(instrumentIds, si.Id)
		table = append(table, fmt.Sprintf("|%s|%s|%s|", si.Name, si.FactoryNumber, si.Person))
	}

	channelId, err := s.channel.Create(ctx, notification.UserId)
	if err != nil {
		return err
	}

	message := strings.Join(table, "\n")

	if notification.Message != "" {
		message = fmt.Sprintf("#### %s\n%s", notification.Message, message)
	}

	// postId := uuid.New()

	post := &model.Post{
		ChannelId: channelId,
		Message:   message,
	}

	if notification.Type == "receiving" {
		url := s.AppUrl + constants.ApiReceiving

		logger.Debug("app url ", url)

		action := &model.PostAction{
			Id:    "all",
			Name:  "Получено",
			Style: "primary",
			Integration: &model.PostActionIntegration{
				URL: url,
				Context: map[string]interface{}{
					"type":          "receiving",
					"status":        notification.Status,
					"instrumentIds": instrumentIds,
				},
			},
		}

		attachment := &model.SlackAttachment{
			Actions: []*model.PostAction{action},
		}

		post.AddProp("attachments", []*model.SlackAttachment{attachment})
	}

	_, _, err = s.MostClient.CreatePost(post)
	if err != nil {
		return fmt.Errorf("failed to create post. error: %w", err)
	}

	// a := post.Attachments()
	// logger.Debug(a[0].Actions[0])
	// Integration в a[0].Actions[0] почему-то нет

	// post, _, err = s.MostClient.GetPost(post.Id, "")
	// if err != nil {
	// 	logger.Error(err)
	// }
	// a := post.Attachments()
	// logger.Debug(a[0].Actions[0])
	// Integration в a[0].Actions[0] почему-то нет

	// а если прописывать в позиции id поста в котором они отправлены

	// s.MostClient.SearchPostsWithParams()

	return nil
}

// send command | список команд на которые реагирует бот
func (s *MessageService) SendCommand(ctx context.Context, channelId string) error {

	//// команды -> вывести список си, получены все си

	commands := []string{
		"***Список*** - вывести список инструментов",
		"***Команды*** - показать все команды",
		"***Ссылка*** - получить ссылку на сервис",
	}

	post := &model.Post{
		ChannelId: channelId,
		Message:   strings.Join(commands, "\n"),
	}

	_, _, err := s.MostClient.CreatePost(post)
	if err != nil {
		return fmt.Errorf("failed to create post. error: %w", err)
	}

	return fmt.Errorf("not implemented")
}

// maybe send link to service
