package repositories

import (
	"go-admin/app/models"
	"go-admin/app/schema"
)

type IUserRepository interface {
	Login(item *schema.Login) (*models.User, error)
	Register(item *schema.Register) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
}
