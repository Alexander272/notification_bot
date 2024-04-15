package models

import "github.com/mattermost/mattermost-server/v6/model"

type PostAction struct {
	UserId      string            `json:"user_id"`
	UserName    string            `json:"user_name"`
	ChannelId   string            `json:"channel_id"`
	ChannelName string            `json:"channel_name"`
	TeamId      string            `json:"team_id"`
	TeamName    string            `json:"team_domain"`
	PostId      string            `json:"post_id"`
	TriggerId   string            `json:"trigger_id" binding:"required"`
	Type        string            `json:"type"`
	DataSource  string            `json:"data_source"`
	Context     PostActionContext `json:"context,omitempty" binding:"required"`
}

type PostActionContext struct {
	Url         string                `json:"url"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	CallbackId  string                `json:"callbackId"`
	SubmitLabel string                `json:"submitLabel"`
	State       string                `json:"state"`
	Fields      []model.DialogElement `json:"fields"`
}
