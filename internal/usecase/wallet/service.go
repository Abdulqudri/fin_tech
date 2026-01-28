package wallet

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/domain/user"
	"github.com/Abdulqudri/fintech/internal/domain/wallet"
	"github.com/google/uuid"
)

type WalletService struct {
	wallet_repo wallet.WalletRepository
	user_repo   user.UserRepository
}

func NewService(wallet_repo wallet.WalletRepository, user_repo user.UserRepository) *WalletService {
	return &WalletService{wallet_repo: wallet_repo, user_repo: user_repo}
}

func (s *WalletService) GetByUserId(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	_, err := s.user_repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	wallet, err := s.wallet_repo.GetByUserId(ctx, id)
	return wallet, err
}

func (s *WalletService) GetById(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	return s.wallet_repo.GetById(ctx, id)
}
