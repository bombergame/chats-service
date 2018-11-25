package rest

import "github.com/satori/go.uuid"

type CommandRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type CommandResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MakePrivateChatRequestData struct {
	ProfileID1 int64 `json:"profile_id_1"`
	ProfileID2 int64 `json:"profile_id_2"`
}

type MakePrivateChatResponseData struct {
	ChatID uuid.UUID `json:"chat_id"`
}
