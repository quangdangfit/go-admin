package impl

import (
	"errors"

	"github.com/jinzhu/copier"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
)

type RoleRepo struct {
	db dbs.IDatabase
}

func NewRoleRepository(db dbs.IDatabase) repositories.IRoleRepository {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if r.db.GetInstance().Where("name = ? ", name).First(&role).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &role, nil
}

func (r *RoleRepo) CreateRole(req *schema.RoleBodyParam) (*models.Role, error) {
	var role models.Role
	copier.Copy(&role, &req)

	if err := r.db.GetInstance().Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
