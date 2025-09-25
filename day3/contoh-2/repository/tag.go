package repository

import (
	"contoh-2/model"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (r *TagRepository) CreateBatch(tags []model.Tag) ([]model.Tag, error) {
	tx := r.db.Create(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

func (r *TagRepository) GetByBlogID(blogID int) ([]model.Tag, error) {
	var tags []model.Tag
	tx := r.db.Where("blog_id = ?", blogID).First(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

func (r *TagRepository) Delete(id int) error {
	tx := r.db.Where("id = ?", id).Delete(model.Tag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
