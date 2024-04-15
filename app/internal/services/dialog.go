package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexander272/si_bot/internal/models"
	"github.com/mattermost/mattermost-server/v6/model"
)

type DialogService struct {
	MostClient *model.Client4
}

func NewDialogService(mostClient *model.Client4) *DialogService {
	return &DialogService{
		MostClient: mostClient,
	}
}

type Dialog interface {
	Open(ctx context.Context, action models.PostAction) error
}

func (s *DialogService) Open(ctx context.Context, action models.PostAction) error {
	state := []string{
		fmt.Sprintf("PostId:%s", action.PostId),
		action.Context.State,
	}

	dialogData := model.Dialog{
		CallbackId:       action.Context.CallbackId,
		Title:            action.Context.Title,
		IntroductionText: action.Context.Description,
		Elements:         action.Context.Fields,
		State:            strings.Join(state, "&"),
	}

	dialog := model.OpenDialogRequest{
		TriggerId: action.TriggerId,
		URL:       action.Context.Url,
		Dialog:    dialogData,
	}

	_, err := s.MostClient.OpenInteractiveDialog(dialog)
	if err != nil {
		return fmt.Errorf("failed to open dialog. error: %w", err)
	}

	return nil
}
