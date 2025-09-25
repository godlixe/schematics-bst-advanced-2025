package dto

type TagDTO struct {
	Name string `json:"name" binding:"required"`
}

type BatchCreateTagsDTO struct {
	Tags []TagDTO `json:"tags"`
}
