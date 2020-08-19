package app

import (
	"github.com/dgrijalva/jwt-go"
	"go-admin/config"

	"go-admin/pkg/errors"
	jwtAuth "go-admin/pkg/jwt"
)

func InitAuth() (jwtAuth.IJWTAuth, error) {

	cfg := config.Config.JWTAuth
	var opts []jwtAuth.Option
	//access token
	opts = append(opts, jwtAuth.WithKeyFunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(cfg.SigningKey), nil
	}))
	opts = append(opts, jwtAuth.WithExpired(cfg.Expired))
	opts = append(opts, jwtAuth.WithSigningKey([]byte(cfg.SigningKey)))
	//refresh token
	opts = append(opts, jwtAuth.WithKeyFuncRefresh(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(cfg.SigningRefreshKey), nil
	}))
	opts = append(opts, jwtAuth.WithExpiredRefresh(cfg.ExpiredRefreshToken))
	opts = append(opts, jwtAuth.WithSigningKeyRefresh([]byte(cfg.SigningRefreshKey)))
	return jwtAuth.NewJWTAuth(opts...), nil
}
