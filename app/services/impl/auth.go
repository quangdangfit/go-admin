package impl

import (
	"context"

	jwtMiddle "go-admin/app/middleware/jwt"
	"go-admin/app/models"
	"go-admin/app/repositories"
	"go-admin/app/schema"
	"go-admin/app/services"
	"go-admin/pkg/jwt"
)

type AuthService struct {
	jwt      jwt.IJWTAuth
	userRepo repositories.IUserRepository
	roleRepo repositories.IRoleRepository
}

func NewAuthService(jwt jwt.IJWTAuth, user repositories.IUserRepository, role repositories.IRoleRepository) services.IAuthService {
	return &AuthService{
		jwt:      jwt,
		userRepo: user,
		roleRepo: role,
	}
}

func (a *AuthService) Login(ctx context.Context, item *schema.Login) (*models.User, string, error) {
	user, err := a.userRepo.Login(item)
	if err != nil {
		return nil, "", err
	}

	token, _ := a.jwt.GenerateToken(user.ID)
	return user, token.GetAccessToken(), nil
}

func (u *AuthService) Register(ctx context.Context, item *schema.Register) (*models.User, string, error) {
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
