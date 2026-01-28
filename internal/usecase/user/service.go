package user

import (
	"context"

	userdomain "github.com/Abdulqudri/fintech/internal/domain/user"
	"github.com/Abdulqudri/fintech/internal/domain/wallet"
	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	pass "github.com/Abdulqudri/fintech/internal/utils/password"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	user_repo userdomain.UserRepository
	waller_repo wallet.WalletRepository
	db *gorm.DB
}

func NewService(user_repo userdomain.UserRepository, waller_repo wallet.WalletRepository) *UserService {
	return &UserService{user_repo: user_repo, waller_repo: waller_repo}
}
func (s *UserService) CreateUser(ctx context.Context, u *userdomain.User, password string) error {
	tx, err := db.BeginTx(ctx, s.db)
	if err != nil {
		return err
	}

	defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

	user := userdomain.User{
		ID: uuid.New(),
		FullName: u.FullName,
		Email:    u.Email,
		Status:   userdomain.StatusActive,		
	}
	hashedPassword, err := pass.Hash(password)
	if err != nil {
		return err
	}

	if err := s.user_repo.CreateTx(ctx, tx, &user, hashedPassword); err != nil {
		return err
	}
	wallet := wallet.Wallet{
		ID: uuid.New(),
		UserID:   user.ID,
		Currency: "NGN",
		Status:   wallet.StatusActive,
	}

	if err := s.waller_repo.CreateTx(ctx, tx, &wallet); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
        return err
    }
	return nil

}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	user, err := s.user_repo.GetByEmail(ctx, email)
	return user, err

}
func (s *UserService) GetById(ctx context.Context, id uuid.UUID) (*userdomain.User, error) {
	user, err := s.user_repo.GetById(ctx, id)
	return user, err

}
