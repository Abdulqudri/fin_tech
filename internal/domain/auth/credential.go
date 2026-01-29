package auth

import "github.com/google/uuid"

type Credential struct {
	ID           uuid.UUID
	HashPassword string
}