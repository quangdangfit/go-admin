package repositories

import (
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	List(queryParam *schema.UserQueryParam) (*[]models.User, error)
	Login(item *schema.LoginBodyParam) (*models.User, error)
	RemoveToken(userId string) (*models.User, error)
	Update(userId string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error)
}
