package user

import (
	"context"

	"github.com/google/uuid"
)

// Repository allows to get/save events from/to event store
type Repository interface {
	Save(ctx context.Context, u *User) error
	Get(id uuid.UUID) *User
}
