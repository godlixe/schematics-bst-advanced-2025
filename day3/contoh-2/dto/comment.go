package dto

type CreateCommentDTO struct {
	BlogID  int    `json:"blog_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
