package service

import (
	"contoh-2/model"
	"errors"
)

type BlogRepository interface {
	Create(blog *model.Blog) (*model.Blog, error)
	GetByID(id int) (*model.Blog, error)
	GetAll() ([]model.Blog, error)
	Update(blog *model.Blog) (*model.Blog, error)
	Delete(id int) error
}

type blogService struct {
	blogRepository BlogRepository
	userRepository UserRepository
}

func NewBlogService(userRepository UserRepository, blogRepository BlogRepository) *blogService {
	return &blogService{
		userRepository: userRepository,
		blogRepository: blogRepository,
	}
}

func (s *blogService) Create(blog *model.Blog) (*model.Blog, error) {
	user, err := s.userRepository.GetByID(blog.UserID)
	if err != nil {
		return nil, err
	}

	blog.Author = user.Name

	return s.blogRepository.Create(blog)
}

func (s *blogService) GetByID(id int) (*model.Blog, error) {
	return s.blogRepository.GetByID(id)
}

func (s *blogService) GetAll() ([]model.Blog, error) {
	return s.blogRepository.GetAll()
}

func (s *blogService) Update(blog *model.Blog) (*model.Blog, error) {
	existing, err := s.blogRepository.GetByID(blog.ID)
	if err != nil {
		return nil, err
	}

	if existing.UserID != blog.UserID {
		return nil, errors.New("unauthorized action")
	}
	return s.blogRepository.Update(existing)
}

func (s *blogService) Delete(id int) error {
	return s.blogRepository.Delete(id)
}
