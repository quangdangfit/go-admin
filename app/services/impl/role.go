package impl

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/app/services"
)

type RoleService struct {
	repo repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) services.IRoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error) {
	var role models.Role
	copier.Copy(&role, &item)
	err := r.repo.Create(&role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
