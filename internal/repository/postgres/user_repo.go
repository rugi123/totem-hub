package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/chirp/internal/config"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepo(ctx context.Context, cfg config.Postgres) (*ChatRepository, error) {
	conn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, cfg.DBName, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(ctx, conn)
	return &ChatRepository{
		Pool: pool,
	}, err
}
