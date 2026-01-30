package auth

import (
	"context"
	"errors"
	"time"

	authdomain "github.com/Abdulqudri/fintech/internal/domain/auth"
	"github.com/Abdulqudri/fintech/internal/utils/password"
	"github.com/google/uuid"
)

var (
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
)

type AuthService struct {
	user_repo authdomain.AuthRepository
	refresh_repo authdomain.RefreshTokenRepository
	refreshExpiry time.Duration
}

func NewService(repo authdomain.AuthRepository, refreshExpiry time.Duration) *AuthService {
	return &AuthService{user_repo: repo, refreshExpiry: refreshExpiry}
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
		ExpiresAt: time.Now().Add(s.refreshExpiry),
	}

	if err := s.refresh_repo.Save(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshTokenID uuid.UUID) (string, *authdomain.RefreshToken, error) {
	token, err := s.refresh_repo.Find(ctx, refreshTokenID)
	if err != nil {
		return "", nil, ErrInvalidRefreshToken
	}
	if time.Now().After(token.ExpiresAt) {
		_ = s.refresh_repo.Delete(ctx, refreshTokenID)
		return "", nil, ErrInvalidRefreshToken
	}

	//give new token and delete old one
	if err := s.refresh_repo.Delete(ctx, refreshTokenID); err != nil {
		return "", nil, err
	}
	newRefreshToken, err := s.IssueRefreshToken(ctx, token.UserID)
	if err != nil {
		return "", nil, err
	}
	return token.UserID.String(), newRefreshToken, nil

}

func (s *AuthService) RevokeRefreshToken(ctx context.Context, refreshTokenID uuid.UUID) error {
	return s.refresh_repo.Delete(ctx, refreshTokenID)
}
