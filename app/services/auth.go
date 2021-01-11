package services

import (
	"context"

	"github.com/quangdangfit/go-admin/app/schema"
)

type IAuthService interface {
	Login(ctx context.Context, bodyParam *schema.LoginBodyParam) (*schema.UserTokenInfo, error)
	Register(ctx context.Context, bodyParam *schema.RegisterBodyParam) (*schema.UserTokenInfo, error)
	Refresh(ctx context.Context, bodyParam *schema.RefreshBodyParam) (*schema.UserTokenInfo, error)
	Logout(ctx context.Context) error
}
