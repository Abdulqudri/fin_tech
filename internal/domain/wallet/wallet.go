package wallet

import "github.com/google/uuid"

type Wallet struct {
	ID       uuid.UUID
	UserID	 uuid.UUID
	Currency string
	Status   Status
}
type Status string
const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBlocked  Status = "blocked"
)