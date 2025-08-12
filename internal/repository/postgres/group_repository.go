package postgres

import "github.com/rugi123/chirp/pkg/database"

type GroupRepository struct {
	*ChatRepository
}

func NewGroupRepository(db *database.Postgres) *GroupRepository {
	return &GroupRepository{ChatRepository: NewChatRepository(db)}
}
