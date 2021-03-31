package jwt

import (
	"encoding/json"
)

// TokenInfo interface
type TokenInfo interface {
	GetAccessToken() string
	GetRefreshToken() string
	GetTokenType() string
	EncodeToJSON() ([]byte, error)
}

type tokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// GetAccessToken return access token
func (t *tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

// GetRefreshToken return refresh token
func (t *tokenInfo) GetRefreshToken() string {
	return t.RefreshToken
}

// GetTokenType return token type
func (t *tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t *tokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}
