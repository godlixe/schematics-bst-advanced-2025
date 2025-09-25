package repository

import (
	"contoh-2/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment *model.Comment) (*model.Comment, error) {
	tx := r.db.Create(&comment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comment, nil
}

func (r *CommentRepository) GetByBlogID(blogID int) ([]model.Comment, error) {
	var comments []model.Comment
	tx := r.db.Where("blog_id = ?", blogID).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

func (r *CommentRepository) GetByUserID(userID int) ([]model.Comment, error) {
	var comments []model.Comment
	tx := r.db.Where("user_id = ?", userID).Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

func (r *CommentRepository) Update(comment *model.Comment) (*model.Comment, error) {
	tx := r.db.Where("id = ?", comment.ID).Updates(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comment, nil
}

func (r *CommentRepository) Delete(id int) error {
	tx := r.db.Where("id = ?", id).Delete(model.Comment{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
