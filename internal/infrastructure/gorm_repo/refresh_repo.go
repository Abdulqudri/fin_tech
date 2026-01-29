package gormrepo

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/domain/auth"
	"github.com/Abdulqudri/fintech/internal/infrastructure/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository (db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Save(ctx context.Context, token *auth.RefreshToken) error {
	model := models.RefreshToken{
		ID: token.ID,
		UserID: token.UserID,
		ExpiresAt: token.ExpiresAt,
	}
	return r.db.WithContext(ctx).Create(&model).Error
}
func (r *RefreshTokenRepository) Find(ctx context.Context, id uuid.UUID) (*auth.RefreshToken, error) {
	var model models.RefreshToken
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	refresh_token := auth.RefreshToken{
		ID: model.ID,
		UserID: model.UserID,
		ExpiresAt: model.ExpiresAt,
	}
	return &refresh_token, nil
}
func (r *RefreshTokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.RefreshToken{}, "id = ?", id).Error
}