package user

import (
	"context"
	"fmt"

	"github.com/rugi123/totem-hub/internal/domain/entity"
	"github.com/rugi123/totem-hub/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.IDRequest, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("gen hash error: %w", err)
	}

	user := entity.NewUser(req.Name, req.Email, string(hash))

	err = u.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}

	return &dto.IDRequest{
		ID: user.ID,
	}, nil
}
