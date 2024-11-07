package auth

import "energia/entities"

type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
}

func FromEntities(user entities.User) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
}
