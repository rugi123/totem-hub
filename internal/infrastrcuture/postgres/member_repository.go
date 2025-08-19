package postgres

import (
	"context"
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

func (r *MemberRepository) GetMembers(ctx context.Context, user_id uuid.UUID) ([]entity.Member, error) {
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

func (r *MemberRepository) GetMember(ctx context.Context, user_id uuid.UUID) (*entity.Member, error) {
	query := `
		SELECT * FROM chat_members
		WHERE user_id = $1`
	var member entity.Member
	err := r.db.Pool.QueryRow(ctx, query, user_id).Scan(&member.ID, &member.UserID, &member.ChatID, &member.Role, &member.IsMuted)
	return nil, err
}
