package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rugi123/chirp/internal/config"
	"github.com/rugi123/chirp/internal/domain"
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

func (r *MemberRepo) GetMember(ctx context.Context, id int) (*entity.ChatMember, error) {
	query := `
		SELECT id, user_id, chat_id, role FROM chat_members
		FROM chat_members
		WHERE id = $1
		`
	var member entity.ChatMember
	err := r.pool.QueryRow(ctx, query, id).Scan(&member.ID, &member.UserID, &member.ChatID, &member.Role)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("context canceled: %w", err)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return &member, nil
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
