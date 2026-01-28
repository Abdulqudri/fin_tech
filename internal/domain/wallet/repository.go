package wallet

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/google/uuid"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *Wallet) error
	CreateTx(ctx context.Context, tx db.Tx, wallet *Wallet) error
	GetById(ctx context.Context, id uuid.UUID) (*Wallet, error)
	GetByUserId(ctx context.Context, userId uuid.UUID) (*Wallet, error)
}