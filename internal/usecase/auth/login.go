package usecase

import (
	"context"
	"fmt"

	"github.com/rugi123/chirp/internal/dto"
	"github.com/rugi123/chirp/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Login(ctx context.Context, req dto.LoginRequest) (*dto.IDRequest, error) {
	// валидацию надо вынести в обработчик чтобы сразу отдавать фронтенду ошибку
	if err := validator.Validate(req); err != nil {
		return nil, fmt.Errorf("validate error: %w", err)
	}

	user, err := u.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("check password error: %w", err)
	}
	//тут jwt логика

	return &dto.IDRequest{
		ID: user.ID,
	}, nil
}
