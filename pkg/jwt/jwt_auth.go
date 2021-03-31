package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/quangdangfit/go-admin/pkg/errors"
)

type IJWTAuth interface {
	GenerateToken(userID string) (TokenInfo, error)
	RefreshToken(refreshToken string) (TokenInfo, error)
	ParseUserID(accessToken string, refresh bool) (string, error)
}

const defaultKey = "gin-go"
const defaultRefreshKey = "refresh-gin-go"

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       7200,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyFunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(defaultKey), nil
	},
	keyFuncRefresh: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenInvalid
		}
		return []byte(defaultRefreshKey), nil
	},
	expiredRefresh:    24,
	signingRefreshKey: []byte(defaultKey),
}

func NewJWTAuth(opts ...Option) *JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{
		opts: &o,
	}
}

type JWTAuth struct {
	opts *options
}
type options struct {
	signingMethod     jwt.SigningMethod
	signingKey        interface{}
	keyFunc           jwt.Keyfunc
	expired           int
	tokenType         string
	keyFuncRefresh    jwt.Keyfunc
	expiredRefresh    int
	signingRefreshKey interface{}
}
type Option func(*options)

func WithExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}

func WithSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

func WithExpiredRefresh(expired int) Option {
	return func(o *options) {
		o.expiredRefresh = expired
	}
}

func WithKeyFuncRefresh(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFuncRefresh = keyFunc
	}
}

func WithSigningKeyRefresh(key interface{}) Option {
	return func(o *options) {
		o.signingRefreshKey = key
	}
}

func (a *JWTAuth) GenerateToken(userID string) (TokenInfo, error) {
	accessToken, err := a.generateAccess(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.generateRefresh(userID)
	if err != nil {
		return nil, err
	}
	tokenInfo := &tokenInfo{
		TokenType:    a.opts.tokenType,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenInfo, nil
}

func (a *JWTAuth) parseToken(tokenString string, refresh bool) (*jwt.StandardClaims, error) {
	option := a.opts.keyFunc
	if refresh == true {
		option = a.opts.keyFuncRefresh
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, option)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.ErrTokenInvalid
			} else {
				return nil, errors.ErrTokenInvalid
			}
		}
	} else if !token.Valid {
		return nil, errors.ErrTokenInvalid
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

func (a *JWTAuth) ParseUserID(tokenString string, refresh bool) (string, error) {
	claims, err := a.parseToken(tokenString, refresh)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}

func (a *JWTAuth) RefreshToken(refreshToken string) (TokenInfo, error) {
	userID, err := a.ParseUserID(refreshToken, true)
	if err != nil {
		if err == errors.ErrTokenExpired {
			return a.GenerateToken(userID)
		}
		return nil, err
	}

	accessToken, err := a.generateAccess(userID)
	if err != nil {
		return nil, err
	}

	tokenInfo := &tokenInfo{
		TokenType:    a.opts.tokenType,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenInfo, nil
}

func (a *JWTAuth) generateAccess(userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(a.opts.signingMethod, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Duration(a.opts.expired) * time.Second).Unix(),
		NotBefore: now.Unix(),
		Subject:   userID,
	})
	tokenString, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return "", errors.New("generate token fail")
	}

	return tokenString, nil
}

func (a *JWTAuth) generateRefresh(userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(a.opts.signingMethod, jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Duration(a.opts.expiredRefresh) * time.Hour).Unix(),
		NotBefore: now.Unix(),
		Subject:   userID,
	})
	tokenString, err := token.SignedString(a.opts.signingRefreshKey)
	if err != nil {
		return "", errors.New("generate token fail")
	}

	return tokenString, nil
}
