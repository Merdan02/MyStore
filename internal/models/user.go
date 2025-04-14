package models

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
