package models

import (
	"github.com/satori/go.uuid"
)

//easyjson:json
type Chat struct {
	ChatID  uuid.UUID `json:"chat_id"`
	RoomID  uuid.UUID `json:"room_id"`
	Players []int64   `json:"player_ids"`
}
