package interfaces

import (
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

// IUserRepository interface
type IUserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	List(queryParam *schema.UserQueryParam) (*[]models.User, error)
	Login(item *schema.LoginBodyParam) (*models.User, error)
	RemoveToken(userID string) (*models.User, error)
	Update(userID string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error)
}
