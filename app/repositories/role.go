package repositories

import (
	"github.com/quangdangfit/go-admin/app/models"
)

type IRoleRepository interface {
	GetByName(name string) (*models.Role, error)
	Create(role *models.Role) error
}
