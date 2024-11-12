package entities

import "time"

type Device struct {
	ID        int
	UserID    int
	Name      string
	Power     float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
