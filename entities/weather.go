package entities

import "time"

type Weather struct {
	ID          int
	City        string
	Date        time.Time
	Temperature float32
	Humidity    float32
	Description string
}
