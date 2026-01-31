package user

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *User, hash string) error
	CreateTx(ctx context.Context, tx db.Tx, user *User, hash string) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
}
