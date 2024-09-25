package usecase

import (
	"car-rental-api/internal/domain"
	"car-rental-api/internal/repository"
)

type UserUsecase interface {
	CreateUser(user *domain.User) error
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}

// UserUsecaseImpl implementation
type UserUsecaseImpl struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{UserRepo: userRepo}
}

func (u *UserUsecaseImpl) CreateUser(user *domain.User) error {
	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecaseImpl) GetUserByID(id uint) (*domain.User, error) {
	return u.UserRepo.GetUserByID(id)
}

func (u *UserUsecaseImpl) GetAllUsers() ([]*domain.User, error) {
	return u.UserRepo.GetAllUsers()
}
