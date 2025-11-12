package repositories

import (
	"github.com/qobilovvv/1uchet/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByPhoneNumber(phone_number string) (*models.User, error)
	GetByID(user_id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByPhoneNumber(phone_number string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone_number = ?", phone_number).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepository) GetByID(user_id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, user_id).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
