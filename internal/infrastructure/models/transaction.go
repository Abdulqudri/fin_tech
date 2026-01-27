package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type           string    `gorm:"not null"`
	Status         string    `gorm:"not null"`
	IdempotencyKey string    `gorm:"uniqueIndex;not null"`
	CreatedAt      time.Time
}
