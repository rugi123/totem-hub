package usecase

import (
	"context"
	"fmt"

	"github.com/rugi123/chirp/internal/domain/entity"
	"github.com/rugi123/chirp/internal/dto"
	"github.com/rugi123/chirp/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.IDRequest, error) {
	if err := validator.Validate(req); err != nil {
		return nil, fmt.Errorf("validate error: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("gen hash error: %w", err)
	}

	user := entity.NewUser(req.Name, req.Email, string(hash))

	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return &dto.IDRequest{
		ID: user.ID,
	}, nil
}
