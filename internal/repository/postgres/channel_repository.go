package postgres

import "github.com/rugi123/chirp/pkg/database"

type ChannelRepository struct {
	*ChatRepository
}

func NewChannelRepository(db *database.Postgres) *ChannelRepository {
	return &ChannelRepository{ChatRepository: NewChatRepository(db)}
}
