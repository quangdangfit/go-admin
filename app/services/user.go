package services

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

type IUserService interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, queryParam *schema.UserQueryParam) (*[]models.User, error)
}
