package models

import (
	"github.com/satori/go.uuid"
)

//easyjson:json
type Message struct {
	MessageID uuid.UUID
	ChatID    uuid.UUID
	ProfileID int64
	Text      string
}
