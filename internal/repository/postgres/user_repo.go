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

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(ctx context.Context, cfg config.Postgres) (*UserRepo, error) {
	conn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.SSLMode)
	pool, err := pgxpool.New(ctx, conn)
	return &UserRepo{
		pool: pool,
	}, err
}

func (r *UserRepo) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
		`
	var user entity.User
	err := r.pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
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
func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
		`
	var user entity.User
	err := r.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
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
func (r *UserRepo) CreateUser(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, name, email, password, created_at)
		VALUES($1, $2, $3, $4, $5)
		`
	_, err := r.pool.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return fmt.Errorf("context canceled: %w", err)
		}
		return fmt.Errorf("postgres error: %w", err)
	}
	return nil
}
func (r *UserRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	return nil
}
func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	return nil
}
