package service

import "contoh-2/model"

type TagRepository interface {
	CreateBatch(tags []model.Tag) ([]model.Tag, error)
	GetAll() ([]model.Tag, error)
	GetByBlogID(blogID int) ([]model.Tag, error)
	Delete(id int) error
	AddTagsToBlog(blogID int, tagIDs []int) error
	RemoveTagsFromBlog(blogID int, tagIDs []int) error
}

type tagService struct {
	tagRepository TagRepository
}

func NewTagService(tagRepository TagRepository) *tagService {
	return &tagService{tagRepository: tagRepository}
}

func (s *tagService) CreateBatch(tags []model.Tag) ([]model.Tag, error) {
	return s.tagRepository.CreateBatch(tags)
}

func (s *tagService) GetAll() ([]model.Tag, error) {
	return s.tagRepository.GetAll()
}

func (s *tagService) GetByBlogID(blogID int) ([]model.Tag, error) {
	return s.tagRepository.GetByBlogID(blogID)
}

func (s *tagService) Delete(id int) error {
	return s.tagRepository.Delete(id)
}

func (s *tagService) AddTagsToBlog(blogID int, tagIDs []int) error {
	return s.tagRepository.AddTagsToBlog(blogID, tagIDs)
}

func (s *tagService) RemoveTagsFromBlog(blogID int, tagIDs []int) error {
	return s.tagRepository.RemoveTagsFromBlog(blogID, tagIDs)
}
