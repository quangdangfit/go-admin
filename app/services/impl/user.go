package impl

import (
	"context"

	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/app/services"
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

func (u *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) List(ctx context.Context, queryParam *schema.UserQueryParam) (*[]models.User, error) {
	user, err := u.userRepo.GetUsers(queryParam)
	if err != nil {
		return nil, err
	}

	return user, nil
}
