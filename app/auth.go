package app

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/quangdangfit/go-admin/config"
	"github.com/quangdangfit/go-admin/pkg/errors"
	jwtAuth "github.com/quangdangfit/go-admin/pkg/jwt"
)

func InitAuth() (jwtAuth.IJWTAuth, error) {
	conf := config.Config.JWTAuth
	var opts []jwtAuth.Option
	//access token
	opts = append(opts, jwtAuth.WithKeyFunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(conf.SigningKey), nil
	}))
	if conf.Expired != 0 {
		opts = append(opts, jwtAuth.WithExpired(conf.Expired))
	}
	opts = append(opts, jwtAuth.WithSigningKey([]byte(conf.SigningKey)))

	//refresh token
	opts = append(opts, jwtAuth.WithKeyFuncRefresh(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(conf.SigningRefreshKey), nil
	}))

	if conf.ExpiredRefreshToken != 0 {
		opts = append(opts, jwtAuth.WithExpiredRefresh(conf.ExpiredRefreshToken))
	}
	opts = append(opts, jwtAuth.WithSigningKeyRefresh([]byte(conf.SigningRefreshKey)))
	return jwtAuth.NewJWTAuth(opts...), nil
}
