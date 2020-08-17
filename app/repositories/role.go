package repositories

import (
	"go-admin/app/models"
	"go-admin/app/schema"
)

type IRoleRepository interface {
	CreateRole(req *schema.RoleBodyParam) (*models.Role, error)
}
