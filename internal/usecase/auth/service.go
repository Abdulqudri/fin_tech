package auth

import (
	"context"
	"errors"
	"time"

	authdomain "github.com/Abdulqudri/fintech/internal/domain/auth"
	"github.com/Abdulqudri/fintech/internal/utils/password"
	"github.com/google/uuid"
)

type AuthService struct {
	user_repo authdomain.AuthRepository
	refresh_repo authdomain.RefreshTokenRepository
}

func NewService(repo authdomain.AuthRepository) *AuthService {
	return &AuthService{user_repo: repo}
}

func (s *AuthService) Login(ctx context.Context, email string, pass string) (*authdomain.Session, error) {
	cred, err := s.user_repo.GetCredentialByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !password.Compare(cred.HashPassword, pass) {
		return nil, errors.New("Invalid Credentials")
	}
	return &authdomain.Session{
		UserID: cred.ID,
	}, nil
}

func (s *AuthService) IssueRefreshToken(
	ctx context.Context,
	userID uuid.UUID,
) (*authdomain.RefreshToken, error) {

	token := &authdomain.RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.refresh_repo.Save(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}
