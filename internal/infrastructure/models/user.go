package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	FullName  string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time
}
