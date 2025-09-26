package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	BlogID    int       `json:"blog_id"`
	Blog      *Blog     `json:"blog,omitempty" gorm:"foreignKey:BlogID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
