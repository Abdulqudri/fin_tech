package wallet

import (
	"context"

	"github.com/google/uuid"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *Wallet) error
	GetById(ctx context.Context, id uuid.UUID) (*Wallet, error)
}