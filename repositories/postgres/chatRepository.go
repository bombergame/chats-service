package postgres

import (
	"github.com/bombergame/common/errs"
	"github.com/satori/go.uuid"
)

type ChatRepository struct {
	conn *Connection
}

func NewProfileRepository(conn *Connection) *ChatRepository {
	return &ChatRepository{
		conn: conn,
	}
}

func (r *ChatRepository) StartPrivateChat(firstProfileID, secondProfileID int64) (uuid.UUID, error) {
	query := `SELECT * FROM create_or_get_profile($1,$2);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return uuid.Nil, errs.NewServiceError(err)
	}

	row := statement.QueryRow(firstProfileID, secondProfileID)

	var chatID uuid.UUID
	if err := row.Scan(&chatID); err != nil {
		return uuid.Nil, wrapError(err)
	}

	return chatID, nil
}

func wrapError(err error) error {
	return errs.NewServiceError(err)
}
