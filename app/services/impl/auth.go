package impl

import (
	"context"

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

func NewAuthService(jwt jwt.IJWTAuth, user repositories.IUserRepository,
	role repositories.IRoleRepository) services.IAuthService {
	return &AuthService{
		jwt:      jwt,
		userRepo: user,
		roleRepo: role,
	}
}

func (a *AuthService) Login(ctx context.Context, item *schema.LoginBodyParam) (*schema.UserTokenInfo, error) {
	user, err := a.userRepo.Login(item)
	if err != nil {
		return nil, err
	}

	token, _ := a.jwt.GenerateToken(user.ID)
	values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
	_, err = a.userRepo.Update(user.ID, &values)
	if err != nil {
		return nil, err
	}

	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
		ExpiresAt:    token.GetExpiresAt(),
	}

	return &tokenInfo, nil
}

func (a *AuthService) Register(ctx context.Context, item *schema.RegisterBodyParam) (*schema.UserTokenInfo, error) {
	if item.RoleID == "" {
		role, err := a.roleRepo.GetRoleByName("user")
		if err != nil {
			return nil, err
		}

		item.RoleID = role.ID
	}

	user, err := a.userRepo.Register(item)
	if err != nil {
		return nil, err
	}

	token, _ := a.jwt.GenerateToken(user.ID)
	values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
	_, err = a.userRepo.Update(user.ID, &values)
	if err != nil {
		return nil, err
	}

	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
		ExpiresAt:    token.GetExpiresAt(),
	}

	return &tokenInfo, nil
}
