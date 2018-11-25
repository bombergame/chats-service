package repositories

import (
	"github.com/satori/go.uuid"
)

type ChatRepository interface {
	StartPrivateChat(firstProfileID, secondProfileID int64) (uuid.UUID, error)
}
