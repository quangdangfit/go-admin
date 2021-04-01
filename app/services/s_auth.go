package services

import (
	"context"

	"github.com/jinzhu/copier"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/app"
	"github.com/quangdangfit/go-admin/pkg/errors"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

// AuthService authentication service
type AuthService struct {
	jwt      jwt.IJWTAuth
	userRepo interfaces.IUserRepository
	roleRepo interfaces.IRoleRepository
}

// NewAuthService return new IAuthService interface
func NewAuthService(jwt jwt.IJWTAuth, user interfaces.IUserRepository,
	role interfaces.IRoleRepository) interfaces.IAuthService {
	return &AuthService{
		jwt:      jwt,
		userRepo: user,
		roleRepo: role,
	}
}

// Login handle user login
func (a *AuthService) Login(ctx context.Context, bodyParam *schema.LoginBodyParam) (*schema.UserTokenInfo, error) {
	user, err := a.userRepo.Login(bodyParam)
	if err != nil {
		return nil, err
	}

	token, err := a.jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
	_, err = a.userRepo.Update(user.ID, &values)
	if err != nil {
		return nil, err
	}

	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
	}

	return &tokenInfo, nil
}

// Register register user
func (a *AuthService) Register(ctx context.Context, param *schema.RegisterBodyParam) (*schema.UserTokenInfo, error) {
	if param.RoleID == "" {
		role, err := a.roleRepo.GetByName("user")
		if err != nil {
			return nil, err
		}

		param.RoleID = role.ID
	}

	var user models.User
	copier.Copy(&user, &param)
	err := a.userRepo.Create(&user)
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
	}

	return &tokenInfo, nil
}

// Refresh refresh token for user
func (a *AuthService) Refresh(ctx context.Context, bodyParam *schema.RefreshBodyParam) (*schema.UserTokenInfo, error) {
	user, err := a.userRepo.GetUserByToken(bodyParam.RefreshToken)
	if err != nil {
		return nil, errors.ErrorTokenInvalid.New()
	}

	token, err := a.jwt.RefreshToken(bodyParam.RefreshToken)
	if err != nil {
		return nil, err
	}

	if token.GetRefreshToken() != user.RefreshToken {
		values := schema.UserUpdateBodyParam{RefreshToken: token.GetRefreshToken()}
		_, err = a.userRepo.Update(user.ID, &values)
	}

	tokenInfo := schema.UserTokenInfo{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		TokenType:    token.GetTokenType(),
	}

	return &tokenInfo, nil
}

// Logout logout user
func (a *AuthService) Logout(ctx context.Context) error {
	_, err := a.userRepo.RemoveToken(app.GetUserID(ctx))
	if err != nil {
		return err
	}

	return nil
}
