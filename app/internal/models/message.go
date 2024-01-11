package models

type Notification struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
	SI      []SI   `json:"si"`
}
