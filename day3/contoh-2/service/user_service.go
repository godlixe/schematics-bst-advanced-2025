package service

import (
	"contoh-2/model"
	"contoh-2/utils/auth"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
}

type JWTService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (int, error)
}

type userService struct {
	jwtService     JWTService
	userRepository UserRepository
}

func NewUserService(jwtService JWTService, userRepository UserRepository) *userService {
	return &userService{
		jwtService:     jwtService,
		userRepository: userRepository,
	}
}

func (s *userService) Register(user *model.User) (*model.User, error) {
	// check if user with the same email exists
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, errors.New("another user with the same email exists")
	}

	// hash user password
	passHash, err := auth.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = passHash

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) Login(user *model.User) (string, error) {
	existing, err := s.userRepository.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if existing == nil {
		return "", errors.New("user does not exist")
	}

	res, err := auth.ComparePassword(existing.PasswordHash, []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !res {
		return "", errors.New("password does not match")
	}

	// return token
	token, err := s.jwtService.GenerateToken(existing.ID)
	if err != nil {
		return "", err
	}

	return token, err
}

func (s *userService) GetByID(id int) (*model.User, error) {
	return s.userRepository.GetByID(id)
}
func (s *userService) Update(user *model.User) (*model.User, error) {
	// hash user password
	passHash, err := auth.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = passHash

	err = s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
