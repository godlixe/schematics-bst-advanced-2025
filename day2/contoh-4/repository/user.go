package repository

import (
	"errors"

	"contoh-3/model"
)

type UserRepository struct {
	seqID int
	data  []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{data: make([]model.User, 0)}
}

func (r *UserRepository) Create(blog model.User) (model.User, error) {
	r.seqID++
	blog.ID = r.seqID
	r.data = append(r.data, blog)
	return blog, nil
}

func (r *UserRepository) GetByEmail(email string) (model.User, error) {

	for _, b := range r.data {
		if b.Email == email {
			return b, nil
		}
	}
	return model.User{}, errors.New("user not found")
}
