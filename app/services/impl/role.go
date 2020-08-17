package impl

import (
	"context"

	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/app/services"
)

type RoleService struct {
	repo repositories.IRoleRepository
}

func NewRoleService(repo repositories.IRoleRepository) services.IRoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error) {
	role, err := r.repo.CreateRole(item)
	if err != nil {
		return nil, err
	}

	return role, nil
}
