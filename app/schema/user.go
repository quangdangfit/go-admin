package schema

type User struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Extra    interface{} `json:"extra,omitempty"`
}

type Register struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   string `json:"role_id"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
