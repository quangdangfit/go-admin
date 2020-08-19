package schema

type User struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Extra    interface{} `json:"extra,omitempty"`
}

type RegisterBodyParam struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   string `json:"role_id"`
}

type LoginBodyParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserQueryParam struct {
	Username string `json:"username,omitempty" form:"username,omitempty"`
	Email    string `json:"email,omitempty" form:"email,omitempty"`
}

type UserTokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
}

type UserUpdateBodyParam struct {
	Password     string `json:"password,omitempty"`
	RoleID       string `json:"role_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
