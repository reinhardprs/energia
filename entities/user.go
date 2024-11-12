package entities

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
}
