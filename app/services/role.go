package services

import (
	"context"

	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
)

type IRoleService interface {
	CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error)
}

type role struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) IRoleService {
	return &role{repo: repo}
}

func (r *role) CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error) {
	role, err := r.repo.CreateRole(item)
	if err != nil {
		return nil, err
	}

	return role, nil
}
