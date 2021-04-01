package interfaces

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
)

// IUserService interface
type IUserService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, param *schema.UserQueryParam) (*[]models.User, error)
}
