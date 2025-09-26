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
	tx := r.db.Table("tags").Select("tags.*").Joins("JOIN blogs_tags bt ON bt.tag_id = tags.id").Where("bt.blog_id = ?", blogID).Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

func (r *TagRepository) GetAll() ([]model.Tag, error) {
	var tags []model.Tag
	tx := r.db.Find(&tags)
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

func (r *TagRepository) AddTagsToBlog(blogID int, tagIDs []int) error {
	var blog model.Blog
	if err := r.db.First(&blog, blogID).Error; err != nil {
		return err
	}
	var tags []model.Tag
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}
	if err := r.db.Model(&blog).Association("Tags").Append(&tags); err != nil {
		return err
	}
	return nil
}

func (r *TagRepository) RemoveTagsFromBlog(blogID int, tagIDs []int) error {
	var blog model.Blog
	if err := r.db.First(&blog, blogID).Error; err != nil {
		return err
	}
	var tags []model.Tag
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}
	if err := r.db.Model(&blog).Association("Tags").Delete(&tags); err != nil {
		return err
	}
	return nil
}
