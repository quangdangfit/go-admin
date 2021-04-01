package interfaces

import (
	"github.com/quangdangfit/go-admin/app/models"
)

// IRoleRepository interface
type IRoleRepository interface {
	GetByName(name string) (*models.Role, error)
	Create(role *models.Role) error
}
