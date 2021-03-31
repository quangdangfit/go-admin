package services

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

type RoleService struct {
	repo interfaces.IRoleRepository
}

func NewRoleService(repo interfaces.IRoleRepository) interfaces.IRoleService {
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