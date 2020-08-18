package services

import (
	"context"

	"go-admin/app/models"
	"go-admin/app/schema"
)

type IUserService interface {
	Login(ctx context.Context, item *schema.Login) (*models.User, string, error)
	Register(ctx context.Context, item *schema.Register) (*models.User, string, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}
