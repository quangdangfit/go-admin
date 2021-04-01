package repositories

import (
	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/pkg/errors"
)

// RoleRepo role repository struct
type RoleRepo struct {
	db interfaces.IDatabase
}

// NewRoleRepository return new IRoleRepository interface
func NewRoleRepository(db interfaces.IDatabase) interfaces.IRoleRepository {
	return &RoleRepo{db: db}
}

// Create new role
func (r *RoleRepo) Create(role *models.Role) error {
	if err := r.db.GetInstance().Create(&role).Error; err != nil {
		return errors.Wrap(err, "RoleRepo.Create")
	}
	return nil
}

// GetByName get role by name
func (r *RoleRepo) GetByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.GetInstance().Where("name = ? ", name).First(&role).Error; err != nil {
		return nil, errors.Wrap(err, "RoleRepo.GetByName")
	}
	return &role, nil
}
