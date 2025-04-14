package models

import "time"

type Orders struct {
	ID        int
	UserId    int
	Status    string
	CreatedAt time.Time
}
