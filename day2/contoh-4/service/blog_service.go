package service

import (
	"time"

	"contoh-3/model"
)

type BlogRepository interface {
	Create(blog model.Blog) (model.Blog, error)
	GetByID(id int) (model.Blog, error)
	GetAll() ([]model.Blog, error)
	Update(blog model.Blog) (model.Blog, error)
	Delete(id int) error
}

type blogService struct {
	blogRepository BlogRepository
}

func NewBlogService(blogRepository BlogRepository) *blogService {
	return &blogService{
		blogRepository: blogRepository,
	}
}

func (s *blogService) Create(blog model.Blog) (model.Blog, error) {
	blog.Timestamp = time.Now()
	return s.blogRepository.Create(blog)
}

func (s *blogService) GetByID(id int) (model.Blog, error) {
	return s.blogRepository.GetByID(id)
}

func (s *blogService) GetAll() ([]model.Blog, error) {
	return s.blogRepository.GetAll()
}

func (s *blogService) Update(id int, blog model.Blog) (model.Blog, error) {
	existing, err := s.blogRepository.GetByID(id)
	if err != nil {
		return model.Blog{}, err
	}
	existing.Timestamp = time.Now()
	return s.blogRepository.Update(blog)
}

func (s *blogService) Delete(id int) error {
	return s.blogRepository.Delete(id)
}
