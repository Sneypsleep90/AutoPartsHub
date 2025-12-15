package storage

import (
	"context"

	"github.com/Sneypsleep90/AutoPartsHub/user-service/internal/models"
)

// user
type User interface {
	Create(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
}
