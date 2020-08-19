package jwt

import (
	"encoding/json"
)

type RefreshInfo interface {
	GetAccessToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeToJSON() ([]byte, error)
}
type refreshInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
}

func (t *refreshInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *refreshInfo) GetTokenType() string {
	return t.TokenType
}

func (t *refreshInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t *refreshInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}
