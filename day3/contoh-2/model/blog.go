package model

import "time"

type Blog struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Comments  []Comment `json:"comments,omitempty"`
	Tags      []Tag     `json:"tags,omitempty" gorm:"many2many:blogs_tags;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
