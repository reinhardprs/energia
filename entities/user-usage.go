package entities

import "time"

type UserUsage struct {
	ID          int
	UserID      int
	Date        time.Time
	TotalEnergy float32
	TotalCost   float32
}
