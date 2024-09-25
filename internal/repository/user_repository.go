package repository

import (
	"car-rental-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}

// UserRepositoryImpl implementation
type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.DB.Preload("Role").First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.DB.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
