package repository

import (
	"errors"
	"sort"
	"time"

	"contoh-3/model"
)

type BlogRepository struct {
	seqID int
	data  []model.Blog
}

func NewBlogRepository() *BlogRepository {
	return &BlogRepository{data: make([]model.Blog, 0)}
}

func (r *BlogRepository) Create(blog model.Blog) (model.Blog, error) {
	r.seqID++
	blog.ID = r.seqID
	r.data = append(r.data, blog)
	return blog, nil
}

func (r *BlogRepository) GetByID(id int) (model.Blog, error) {

	for _, b := range r.data {
		if b.ID == id {
			return b, nil
		}
	}
	return model.Blog{}, errors.New("blog not found")
}

func (r *BlogRepository) GetAll() ([]model.Blog, error) {
	result := make([]model.Blog, len(r.data))
	copy(result, r.data)
	sort.Slice(result, func(i, j int) bool { return result[i].Timestamp.Before(result[j].Timestamp) })
	return result, nil
}

func (r *BlogRepository) Update(id int, updated model.Blog) (model.Blog, error) {
	for i := range r.data {
		if r.data[i].ID == id {
			updated.ID = id
			if updated.Timestamp.IsZero() {
				updated.Timestamp = time.Now()
			}
			r.data[i] = updated
			return updated, nil
		}
	}
	return model.Blog{}, errors.New("blog not found")
}

func (r *BlogRepository) Delete(id int) error {
	for i := range r.data {
		if r.data[i].ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return errors.New("blog not found")
}
