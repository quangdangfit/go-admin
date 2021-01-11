package repositories

import (
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

type IRoleRepository interface {
	CreateRole(req *schema.RoleBodyParam) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
}
