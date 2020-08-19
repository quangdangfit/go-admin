package services

import (
	"context"

	"go-admin/app/models"
	"go-admin/app/schema"
)

type IAuthService interface {
	Login(ctx context.Context, item *schema.Login) (*models.User, string, error)
	Register(ctx context.Context, item *schema.Register) (*models.User, string, error)
	//Logout(ctx context.Context) error
}
