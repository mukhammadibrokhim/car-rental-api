package usecase

import (
	"car-rental-api/internal/domain"
	"car-rental-api/internal/repository"
)

type RoleUsecase interface {
	CreateRole(role *domain.Role) error
	GetRoleByID(id uint) (*domain.Role, error)
}

type RoleUsecaseImpl struct {
	RoleRepo repository.RoleRepository
}

func NewRoleUsecase(roleRepo repository.RoleRepository) RoleUsecase {
	return &RoleUsecaseImpl{RoleRepo: roleRepo}
}

func (r *RoleUsecaseImpl) CreateRole(role *domain.Role) error {
	return r.RoleRepo.CreateRole(role)
}

func (r *RoleUsecaseImpl) GetRoleByID(id uint) (*domain.Role, error) {
	return r.RoleRepo.GetRoleByID(id)
}
