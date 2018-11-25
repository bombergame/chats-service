package models

import (
	"github.com/satori/go.uuid"
)

//easyjson:json
type Message struct {
	ID        uuid.UUID
	ProfileID int64
}
