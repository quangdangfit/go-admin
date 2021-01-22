package impl

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/app/services"
	"github.com/quangdangfit/go-admin/config"
	"github.com/quangdangfit/go-admin/pkg/errors"
)

type UserService struct {
	userRepo repositories.IUserRepository
	roleRepo repositories.IRoleRepository
}

func NewUserService(user repositories.IUserRepository, role repositories.IRoleRepository) services.IUserService {
	return &UserService{
		userRepo: user,
		roleRepo: role,
	}
}

func (u *UserService) checkPermission(id string, data map[string]interface{}) bool {
	return data["id"] == id
}

func (u *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "UserService.GetByID")
	}

	return user, nil
}

func (u *UserService) List(ctx context.Context, param *schema.UserQueryParam) (*[]models.User, error) {
	if param.Limit > config.Config.DefaultLimit {
		param.Limit = config.Config.MaxLimit
	} else if param.Limit <= 0 {
		param.Limit = config.Config.DefaultLimit
	}

	user, err := u.userRepo.List(param)
	if err != nil {
		return nil, err
	}

	return user, nil
}
