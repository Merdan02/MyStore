package models

import "time"

type Orders struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
