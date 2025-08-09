package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain/entity"
)

type MemberRepo struct {
	pool *pgxpool.Pool
}

func NewMemberRepo(ctx context.Context, cfg config.Postgres) (*MemberRepo, error) {
	conn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(ctx, conn)
	return &MemberRepo{
		pool: pool,
	}, err
}

func (r *MemberRepo) GetMemberIDs(ctx context.Context, user_id int) ([]uuid.UUID, error) { // всеее  в монг дб
	query := `
		SELECT id FROM chat_members
		WHERE user_id = $1
		`
	rows, err := r.pool.Query(ctx, query, user_id)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	defer rows.Close()

	IDs := make([]uuid.UUID, 0)
	for rows.Next() {
		var id uuid.UUID
		if err = rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		IDs = append(IDs, id)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return IDs, nil
}
func (r *MemberRepo) CreateMember(ctx context.Context, member *entity.ChatMember) error {
	query := `
		INSERT INTO chat_members (id, user_id, chat_id, role)
		VALUES($1, $2, $3, $4)
		`
	_, err := r.pool.Exec(ctx, query, member.ID, member.UserID, member.ChatID, member.Role)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *MemberRepo) UpdateMember(ctx context.Context, member *entity.ChatMember) error {
	return nil
}
func (r *MemberRepo) DeleteMember(ctx context.Context, id int) error {
	return nil
}
