package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type LedgerEntry struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	WalletID      uuid.UUID `gorm:"type:uuid;index;not null"`
	TransactionID uuid.UUID `gorm:"type:uuid;index;not null"`
	Amount        int64     `gorm:"not null"` // signed
	EntryType     string    `gorm:"not null"`
	CreatedAt     time.Time
}
