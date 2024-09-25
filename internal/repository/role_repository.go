package repository

import (
	"car-rental-api/internal/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *domain.Role) error
	GetRoleByID(id uint) (*domain.Role, error)
}

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{DB: db}
}

func (r *RoleRepositoryImpl) CreateRole(role *domain.Role) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepositoryImpl) GetRoleByID(id uint) (*domain.Role, error) {
	var role domain.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}
