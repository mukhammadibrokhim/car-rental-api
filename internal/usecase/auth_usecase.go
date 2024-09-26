package usecase

import (
	"car-rental-api/internal/domain"
	"car-rental-api/internal/domain/payload"
	"car-rental-api/internal/repository"
	"car-rental-api/pkg"
	"errors"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
	Register(request payload.RegisterRequest) (string, error)
}

type AuthUsecaseImpl struct {
	UserRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{UserRepo: userRepo}
}

func (u *AuthUsecaseImpl) Login(email, password string) (string, error) {
	user, err := u.UserRepo.GetUserByUserEmail(email)
	if err != nil || !pkg.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := pkg.GenerateToken(user.ID, user.Role.Title)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecaseImpl) Register(request payload.RegisterRequest) (string, error) {
	hashedPassword, err := pkg.HashPassword(request.Password)
	if err != nil {
		return "", err
	}

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		DOB:      request.Dob,
		Address:  request.Address,
		RoleID:   1,
	}

	if err := u.UserRepo.CreateUser(&user); err != nil {
		return "", err
	}
	token, err := pkg.GenerateToken(user.ID, user.Role.Title)
	if err != nil {
		return "", err
	}
	return token, nil
}
