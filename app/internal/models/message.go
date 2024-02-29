package models

type Notification struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	SI      []SI   `json:"si"`
}
