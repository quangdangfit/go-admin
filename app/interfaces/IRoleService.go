package interfaces

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

// IRoleService interface
type IRoleService interface {
	CreateRole(ctx context.Context, item *schema.RoleBodyParam) (*models.Role, error)
}
