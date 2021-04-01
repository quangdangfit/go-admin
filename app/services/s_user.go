package services

import (
	"context"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/config"
	"github.com/quangdangfit/go-admin/pkg/errors"
)

// UserService user service
type UserService struct {
	userRepo interfaces.IUserRepository
	roleRepo interfaces.IRoleRepository
}

// NewUserService return new IUserService interface
func NewUserService(user interfaces.IUserRepository, role interfaces.IRoleRepository) interfaces.IUserService {
	return &UserService{
		userRepo: user,
		roleRepo: role,
	}
}

func (u *UserService) checkPermission(id string, data map[string]interface{}) bool {
	return data["id"] == id
}

// GetByID get user by id
func (u *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "UserService.GetByID")
	}

	return user, nil
}

// List users by query
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
