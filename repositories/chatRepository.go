package repositories

import (
	"github.com/bombergame/chats-service/models"
	"github.com/satori/go.uuid"
)

type ChatRepository interface {
	CreateChat(chat models.Chat) error
	GetChat(id uuid.UUID) (models.Chat, error)
	UpdateChat(chat models.Chat) error
	DeleteChat(id uuid.UUID) error
}
