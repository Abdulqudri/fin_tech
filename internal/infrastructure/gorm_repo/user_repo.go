package gormrepo

import (
	"context"

	"github.com/Abdulqudri/fintech/internal/domain/user"
	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/Abdulqudri/fintech/internal/infrastructure/models"
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User, hash string) error {

	model := models.User{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Password: hash,
		Status:   string(u.Status),
	}

	return r.db.WithContext(ctx).Create(&model).Error
}

func (r *UserRepository) CreateTx(ctx context.Context, tx db.Tx, u *user.User, hash string) error {

	model := models.User{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
		Password: hash,
		Status:   string(u.Status),
	}

	return tx.DB().WithContext(ctx).Create(&model).Error
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var model models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return &user.User{
		ID:       model.ID,
		FullName: model.FullName,
		Email:    model.Email,
		Status:   user.Status(model.Status),
	}, nil
}

func (r *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var model models.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &user.User{
		ID:       model.ID,
		FullName: model.FullName,
		Email:    model.Email,
		Status:   user.Status(model.Status),
	}, nil
}

