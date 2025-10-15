package repository

import (
	"context"
	"it-backend/internal/model/entity"
)

type AuthRepository interface {
	FindUserByUsername(ctx context.Context, username string) (entity.User, error)
}
