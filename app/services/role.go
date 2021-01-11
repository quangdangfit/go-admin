package services

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

type IRoleService interface {
	CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error)
}
