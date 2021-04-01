package interfaces

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

// IRoleService interface
type IRoleService interface {
	Create(ctx context.Context, item *schema.RoleBodyParams) (*models.Role, error)
}
