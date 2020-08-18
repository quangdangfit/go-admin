package impl

import (
	"context"

	jwtMiddle "go-admin/app/middleware/jwt"
	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/app/services"
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

func (u *UserService) Login(ctx context.Context, item *schema.Login) (*models.User, string, error) {
	user, err := u.userRepo.Login(item)
	if err != nil {
		return nil, "", err
	}

	token := jwtMiddle.GenerateToken(user)
	return user, token, nil
}

func (u *UserService) Register(ctx context.Context, item *schema.Register) (*models.User, string, error) {
	if item.RoleID == "" {
		role, err := u.roleRepo.GetRoleByName("user")
		if err != nil {
			return nil, "", err
		}

		item.RoleID = role.ID
	}

	user, err := u.userRepo.Register(item)
	if err != nil {
		return nil, "", err
	}

	token := jwtMiddle.GenerateToken(user)
	return user, token, nil
}

func (u *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
