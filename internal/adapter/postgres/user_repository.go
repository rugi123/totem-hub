package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rugi123/totem-hub/internal/domain"
	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/pkg/database"
)

type UserRepository struct {
	*PostgresRepository
}

func NewUserRepository(db *database.Postgres) *UserRepository {
	return &UserRepository{PostgresRepository: NewPostgresRepository(db)}
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
		`
	var user entity.User
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("context canceled: %w", err)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return &user, nil
}
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
		`
	var user entity.User
	err := r.db.Pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("context canceled: %w", err)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("postgres error: %w", err)
	}
	return &user, nil
}
func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, name, email, password, created_at)
		VALUES($1, $2, $3, $4, $5)
		`
	_, err := r.db.Pool.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	return nil
}
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
