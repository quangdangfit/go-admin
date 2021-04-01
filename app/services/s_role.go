package services

import (
	"context"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

// RoleService role service
type RoleService struct {
	repo interfaces.IRoleRepository
}

// NewRoleService return new IRoleService interface
func NewRoleService(repo interfaces.IRoleRepository) interfaces.IRoleService {
	return &RoleService{repo: repo}
}

// CreateRole create new role
func (r *RoleService) Create(ctx context.Context, item *schema.RoleBodyParams) (*models.Role, error) {
	var role models.Role
	err := utils.Copy(&role, &item)
	if err != nil {
		return nil, err
	}

	err = r.repo.Create(&role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
