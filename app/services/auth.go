package services

import (
	"context"

	"go-admin/app/schema"
)

type IAuthService interface {
	Login(ctx context.Context, item *schema.LoginBodyParam) (*schema.UserTokenInfo, error)
	Register(ctx context.Context, item *schema.RegisterBodyParam) (*schema.UserTokenInfo, error)
	//Logout(ctx context.Context) error
}
