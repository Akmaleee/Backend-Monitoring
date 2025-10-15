package repository

import (
	"context"
	"errors"
	"it-backend/database"
	"it-backend/internal/model/entity"
)

type AuthRepositoryImpl struct {
	DB database.DatabaseMySQL
}

func NewAuthRepository(db database.DatabaseMySQL) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (r *AuthRepositoryImpl) FindUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var (
		user entity.User
		err  error
	)

	err = r.DB.DBInfra.Preload("RoleUser.Role").Where("username = ?", username).Where("is_active = ?", true).First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
