package impl

import (
	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/pkg/errors"
)

type RoleRepo struct {
	db interfaces.IDatabase
}

func NewRoleRepository(db interfaces.IDatabase) interfaces.IRoleRepository {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) Create(role *models.Role) error {
	if err := r.db.GetInstance().Create(&role).Error; err != nil {
		return errors.Wrap(err, "RoleRepo.Create")
	}
	return nil
}

func (r *RoleRepo) GetByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.GetInstance().Where("name = ? ", name).First(&role).Error; err != nil {
		return nil, errors.Wrap(err, "RoleRepo.GetByName")
	}
	return &role, nil
}
