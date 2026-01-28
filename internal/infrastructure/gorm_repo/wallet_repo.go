package gormrepo

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/domain/wallet"
	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/Abdulqudri/fintech/internal/infrastructure/models"
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(ctx context.Context, w *wallet.Wallet) error {

	model := models.Wallet{
		ID: w.ID,
		UserID: w.UserID,
		Currency: w.Currency,
		Status: string(wallet.StatusActive),
	}
	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *WalletRepository) CreateTx(ctx context.Context, tx db.Tx, w *wallet.Wallet) error {
	model := models.Wallet{
		ID: w.ID,
		UserID: w.UserID,
		Currency: w.Currency,
		Status: string(wallet.StatusActive),
	}
	return tx.DB().WithContext(ctx).Create(&model).Error
}

func (r *WalletRepository) GetById(ctx context.Context, id uuid.UUID) (*wallet.Wallet, error) {
	var model models.Wallet
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &wallet.Wallet{
		ID:       model.ID,
		UserID: model.UserID,
		Currency: model.Currency,
		Status:   wallet.Status(model.Status),
	}, nil
}

func (r *WalletRepository) GetByUserId(ctx context.Context, userId uuid.UUID) (*wallet.Wallet, error) {
	var model models.Wallet

	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&model).Error; err != nil {
		return nil, err
	}
	return &wallet.Wallet{
		ID:       model.ID,
		UserID:   model.UserID,
		Currency: model.Currency,
		Status:   wallet.Status(model.Status),
	}, nil
}