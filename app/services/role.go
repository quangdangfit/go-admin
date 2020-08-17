package services

import (
	"context"

	"go-admin/app/models"
	"go-admin/app/schema"
)

type IRoleService interface {
	CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error)
}
