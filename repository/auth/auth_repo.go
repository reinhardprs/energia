package auth

import (
	"energia/entities"

	"gorm.io/gorm"
)

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Login(user entities.User) (entities.User, error) {
	userDb := FromEntities(user)
	result := authRepo.db.First(&userDb, "email = ?", userDb.Email)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return userDb.ToEntities(), nil
}

func (authRepo AuthRepo) Register(user entities.User) (entities.User, error) {
	userDb := FromEntities(user)
	result := authRepo.db.Create(&userDb)

	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return userDb.ToEntities(), nil
}
