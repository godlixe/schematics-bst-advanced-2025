package model

type BlogsTags struct {
	BlogID int   `json:"blog_id" gorm:"primaryKey;"`
	Blog   *Blog `json:"blog,omitempty" gorm:"foreignKey:BlogID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
	TagID  int   `json:"tag_id" gorm:"primaryKey"`
	Tag    *Tag  `json:"tag,omitempty" gorm:"foreignKey:TagID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;"`
}
