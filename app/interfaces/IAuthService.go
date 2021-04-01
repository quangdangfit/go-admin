package interfaces

import (
	"context"

	"github.com/quangdangfit/go-admin/app/schema"
)

// IAuthService interface
type IAuthService interface {
	Login(ctx context.Context, bodyParam *schema.LoginBodyParams) (*schema.UserTokenInfo, error)
	Register(ctx context.Context, param *schema.RegisterBodyParams) (*schema.UserTokenInfo, error)
	Refresh(ctx context.Context, bodyParam *schema.RefreshBodyParams) (*schema.UserTokenInfo, error)
	Logout(ctx context.Context) error
}
