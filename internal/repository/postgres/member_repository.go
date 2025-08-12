package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/pkg/database"
)

type MemberRepository struct {
	*PostgresRepository
}

func NewMemberRepository(db *database.Postgres) *MemberRepository {
	return &MemberRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *MemberRepository) GetAll(ctx context.Context, user_id uuid.UUID) ([]entity.Member, error) {
	query := `
		SELECT * FROM chat_members
		WHERE user_id = $1
		`
	rows, err := r.db.Pool.Query(ctx, query, user_id)
	if err != nil {
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	defer rows.Close()

	var members []entity.Member
	for rows.Next() {
		var member entity.Member
		if err = rows.Scan(&member.ID, &member.UserID, &member.ChatID, &member.Role, &member.IsMuted); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		members = append(members, member)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}
func (r *MemberRepository) Create(ctx context.Context, member *entity.Member) error {
	query := `
		INSERT INTO chat_members (id, user_id, chat_id, role)
		VALUES($1, $2, $3, $4)
		`
	_, err := r.db.Pool.Exec(ctx, query, member.ID, member.UserID, member.ChatID, member.Role)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *MemberRepository) Update(ctx context.Context, member *entity.Member) error {
	return nil
}
func (r *MemberRepository) Delete(ctx context.Context, member_id uuid.UUID) error {
	return nil
}
