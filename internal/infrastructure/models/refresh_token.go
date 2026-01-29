package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	ExpiresAt time.Time
}