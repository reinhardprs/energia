package user_usage

import (
	"energia/entities"

	"gorm.io/gorm"
)

func NewUserUsageRepo(db *gorm.DB) *UserUsageRepo {
	return &UserUsageRepo{
		db: db,
	}
}

type UserUsageRepo struct {
	db *gorm.DB
}

func (userUsageRepo *UserUsageRepo) Create(userUsage entities.UserUsage) (entities.UserUsage, error) {
	userUsageDb := FromEntities(userUsage)

	result := userUsageRepo.db.Create(&userUsageDb)

	if result.Error != nil {
		return entities.UserUsage{}, result.Error
	}

	return userUsageDb.ToEntities(), nil
}

func (userUsageRepo *UserUsageRepo) FindAll(userID int) ([]entities.UserUsage, error) {
	var userUsagesDb []UserUsage

	// Mengambil data user_usage berdasarkan user_id
	result := userUsageRepo.db.Where("user_id = ?", userID).Find(&userUsagesDb)
	if result.Error != nil {
		return nil, result.Error
	}

	// Mengonversi hasil query ke entitas yang sesuai
	userUsages := make([]entities.UserUsage, len(userUsagesDb))
	for i, userUsageDb := range userUsagesDb {
		userUsages[i] = userUsageDb.ToEntities()
	}

	return userUsages, nil
}
