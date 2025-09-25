package repository

import (
	"contoh-2/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	tx := r.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	tx := r.db.Where("email = ?", email).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}

	// Handle case where record does not exist, this is because gorm returns
	// an error when a record does not exist.
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	fmt.Println(user)
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *UserRepository) Update(user *model.User) error {
	tx := r.db.Where("id = ?", user.ID).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
