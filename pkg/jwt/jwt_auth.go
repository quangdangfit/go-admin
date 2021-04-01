package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/quangdangfit/go-admin/pkg/errors"
)

// IJWTAuth interface
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

// NewJWTAuth return new Auth pointer
func NewJWTAuth(opts ...Option) *Auth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	return &Auth{
		opts: &o,
	}
}

// Auth struct
type Auth struct {
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

// Option jwt option
type Option func(*options)

// WithExpired set expired time
func WithExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// WithKeyFunc set key function
func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}

// WithSigningKey set signing key
func WithSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// WithExpiredRefresh set expired for refresh token
func WithExpiredRefresh(expired int) Option {
	return func(o *options) {
		o.expiredRefresh = expired
	}
}

// WithKeyFuncRefresh set key function for refresh token
func WithKeyFuncRefresh(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFuncRefresh = keyFunc
	}
}

// WithSigningKeyRefresh set signing key for refresh token
func WithSigningKeyRefresh(key interface{}) Option {
	return func(o *options) {
		o.signingRefreshKey = key
	}
}

// GenerateToken return new TokenInfo, generate new access and refresh token
func (a *Auth) GenerateToken(userID string) (TokenInfo, error) {
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

// parseToken parse claims from token
func (a *Auth) parseToken(tokenString string, refresh bool) (*jwt.StandardClaims, error) {
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

// ParseUserID parse user_id from token
func (a *Auth) ParseUserID(tokenString string, refresh bool) (string, error) {
	claims, err := a.parseToken(tokenString, refresh)
	if err != nil {
		return "", err
	}
	return claims.Subject, nil
}

// RefreshToken refresh token, return new TokenInfo
func (a *Auth) RefreshToken(refreshToken string) (TokenInfo, error) {
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

// generateAccess generate access token
func (a *Auth) generateAccess(userID string) (string, error) {
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

// generateRefresh generate refresh token
func (a *Auth) generateRefresh(userID string) (string, error) {
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
