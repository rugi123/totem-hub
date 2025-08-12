package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rugi123/chirp/internal/dto"
	"github.com/rugi123/chirp/internal/transport"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := u.UserRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("get user error: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("check password error: %w", err)
	}

	expirationTime := time.Now().Add(15 + time.Minute)
	claims := &transport.Claims{
		UserID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	fmt.Println("user ", user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenString, err := token.SignedString([]byte(u.Config.App.JWTKey))
	if err != nil {
		return "", fmt.Errorf("create token error: %w", err)
	}
	fmt.Println(tokenString)

	return tokenString, nil
}
