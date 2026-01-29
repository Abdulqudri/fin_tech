package auth

import "github.com/google/uuid"

type Session struct {
	UserID uuid.UUID
}
