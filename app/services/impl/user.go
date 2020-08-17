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
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) services.IUserService {
	return &UserService{repo: repo}
}

func (u *UserService) checkPermission(uuid string, data map[string]interface{}) bool {
	return data["uuid"] == uuid
}

func (u *UserService) Login(ctx context.Context, item *schema.Login) (*models.User, string, error) {
	user, err := u.repo.Login(item)
	if err != nil {
		return nil, "", err
	}

	token := jwtMiddle.GenerateToken(user)
	return user, token, nil
}

func (u *UserService) Register(ctx context.Context, item *schema.Register) (*models.User, string, error) {
	user, err := u.repo.Register(item)
	if err != nil {
		return nil, "", err
	}

	token := jwtMiddle.GenerateToken(user)
	return user, token, nil
}

func (u *UserService) GetUserByID(ctx context.Context, uuid string) (*models.User, error) {
	user, err := u.repo.GetUserByID(uuid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
