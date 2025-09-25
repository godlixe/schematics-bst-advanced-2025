package service

import (
	"contoh-2/model"
	"errors"
)

type CommentRepository interface {
	Create(comment *model.Comment) (*model.Comment, error)
	GetByUserID(userID int) ([]model.Comment, error)
	Update(comment *model.Comment) (*model.Comment, error)
}

type commentService struct {
	commentRepository CommentRepository
}

func NewCommentService(commentRepository CommentRepository) *commentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) Create(comment *model.Comment) (*model.Comment, error) {
	comment, err := s.commentRepository.Create(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) GetByUserID(userID int) ([]model.Comment, error) {
	return s.commentRepository.GetByUserID(userID)
}

func (s *commentService) Update(comment *model.Comment) (*model.Comment, error) {
	existing, err := s.commentRepository.Update(comment)
	if err != nil {
		return nil, err
	}

	if existing.UserID != comment.UserID {
		return nil, errors.New("unauthorized action")
	}
	return s.commentRepository.Update(existing)

}
