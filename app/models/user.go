package models

type User struct {
	Model    `json:"inline"`
	Username string `json:"username" gorm:"unique;not null;index"`
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password string `json:"password"`
	RoleID   string `json:"role_id"`
}
