package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)
type Transfer struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	TransactionID uuid.UUID `gorm:"type:uuid;not null"`
	FromWalletID  uuid.UUID `gorm:"type:uuid;not null"`
	ToWalletID    uuid.UUID `gorm:"type:uuid;not null"`
	Amount        int64     `gorm:"not null"`
	CreatedAt     time.Time
}
