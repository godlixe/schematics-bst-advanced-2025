package repository

import (
	"contoh-2/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) Create(blog *model.Blog) (*model.Blog, error) {
	tx := r.db.Create(&blog)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blog, nil
}

func (r *BlogRepository) GetByID(id int) (*model.Blog, error) {
	var blog model.Blog
	tx := r.db.Where("id = ?", id).First(&blog)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &blog, nil
}

func (r *BlogRepository) GetAll() ([]model.Blog, error) {
	var blogs []model.Blog
	tx := r.db.Preload("Comments").Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

func (r *BlogRepository) Update(blog *model.Blog) (*model.Blog, error) {
	tx := r.db.Where("id = ?", blog.ID).Updates(blog)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blog, nil
}

func (r *BlogRepository) Delete(id int) error {
	tx := r.db.Where("id = ?", id).Delete(model.Blog{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
