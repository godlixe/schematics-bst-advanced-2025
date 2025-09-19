package model

import "time"

type Blog struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Timestamp time.Time `json:"timestamp"`
}
