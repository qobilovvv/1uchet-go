package services

import (
	"errors"

	"github.com/qobilovvv/1uchet/models"
	"github.com/qobilovvv/1uchet/repositories"
)

type UserService interface {
	Create(phone_number string, password string) (*models.User, error)
	GetAll() ([]models.User, error)
	GetByID(user_id uint) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) Create(phone_number string, password string) (*models.User, error) {
	existing, err := s.repo.FindByPhoneNumber(phone_number)
	if err == nil && existing != nil {
		return nil, errors.New("phone number already exists")
	}

	user := &models.User{
		PhoneNumber: phone_number,
		Password:    password,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(user_id uint) (*models.User, error) {
	return s.repo.GetByID(user_id)
}
