package model

import "time"

type User struct {
	ID           int
	Email        string
	Name         string
	PasswordHash string
	Timestamp    time.Time
}
