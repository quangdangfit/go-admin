package impl

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/dbs"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepository() repositories.IRoleRepository {
	return &RoleRepo{db: dbs.Database}
}

func (r *RoleRepo) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if dbs.Database.Where("name = ? ", name).First(&role).RecordNotFound() {
		return nil, errors.New("user not found")
	}

	return &role, nil
}

func (r *RoleRepo) CreateRole(req *schema.RoleBodyParam) (*models.Role, error) {
	var role models.Role
	copier.Copy(&role, &req)

	if err := r.db.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
