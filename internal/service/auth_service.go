package service

import (
	"context"
	"it-backend/internal/model/dto"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
	LoginLDAP(ctx context.Context, request dto.LoginRequest) (string, error)
	CheckLDAP(ctx context.Context, request dto.LoginRequest) error
}
