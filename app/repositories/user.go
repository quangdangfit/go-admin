package repositories

import (
	"go-admin/app/models"
	"go-admin/app/schema"
)

type IUserRepository interface {
	Login(item *schema.LoginBodyParam) (*models.User, error)
	Register(item *schema.RegisterBodyParam) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetUserByToken(token string) (*models.User, error)
	GetUsers(queryParam *schema.UserQueryParam) (*[]models.User, error)
	Update(userId string, bodyParam *schema.UserUpdateBodyParam) (*models.User, error)
}
