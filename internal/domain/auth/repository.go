package auth

import (
	"context"

	"github.com/google/uuid"
)

type AuthRepository interface {
	GetCredentialByEmail(ctx context.Context, email string) (*Credential, error)
}

type RefreshTokenRepository interface {
	Save(ctx context.Context, token *RefreshToken) error
	Find(ctx context.Context, id uuid.UUID) (*RefreshToken, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
