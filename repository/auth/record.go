package auth

import (
	"energia/entities"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `gorm:"unique;type:varchar(255);not null"`
	Password  string    `gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func FromEntities(user entities.User) User {
	return User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func (user User) ToEntities() entities.User {
	return entities.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}
