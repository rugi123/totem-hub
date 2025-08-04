package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/chirp/internal/config"
)

type ChatRepo struct {
	pool *pgxpool.Pool
}

func NewChatRepo(ctx context.Context, cfg config.Postgres) (*ChatRepo, error) {
	conn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(ctx, conn)
	return &ChatRepo{
		pool: pool,
	}, err
}
