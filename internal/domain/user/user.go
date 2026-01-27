package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Status   Status `json:"status"`
}

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBlocked  Status = "blocked"
)
