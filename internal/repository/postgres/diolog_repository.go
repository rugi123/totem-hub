package postgres

import "github.com/rugi123/chirp/pkg/database"

type DiologRepository struct {
	*ChatRepository
}

func NewDiologRepository(db *database.Postgres) *DiologRepository {
	return &DiologRepository{ChatRepository: NewChatRepository(db)}
}
