package dto

type CreateBlogDTO struct {
	Title   string `json:"title,omitempty" binding:"required"`
	Content string `json:"content,omitempty" binding:"required"`
}

type UpdateBlogDTO struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type AddBlogTagsDTO struct {
	TagIDs []int `json:"tag_ids"`
}

type DeleteBlogTagsDTO struct {
	TagIDs []int `json:"tag_ids"`
}
