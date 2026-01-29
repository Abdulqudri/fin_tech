package gormrepo

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/infrastructure/models"
	authdomain "github.com/Abdulqudri/fintech/internal/domain/auth"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}		

func (r *AuthRepository) GetCredentialByEmail(ctx context.Context, email string) (*authdomain.Credential, error) {
	var model models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return &authdomain.Credential{
		ID:           model.ID,
		HashPassword: model.Password,
	}, nil	
}