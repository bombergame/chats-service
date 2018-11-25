package rest

import "github.com/satori/go.uuid"

type AuthID struct {
	Type   string
	AuthID int64 `json:"auth_id"`
}

type CommandRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type CommandResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MakePrivateChatRequestData struct {
	ProfileID int64 `json:"profile_id"`
}

type MakePrivateChatResponseData struct {
	ChatID uuid.UUID `json:"chat_id"`
}
