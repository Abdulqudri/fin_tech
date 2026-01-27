package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	Currency  string    `gorm:"size:3;not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time
}
