package models

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   time.Time
}
