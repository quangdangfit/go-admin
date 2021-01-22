package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/go-admin/pkg/errors"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

type User struct {
	Model        `json:"inline"`
	Username     string `json:"username" gorm:"unique;not null;index"`
	Email        string `json:"email" gorm:"unique;not null;index"`
	Password     string `json:"password" gorm:"not null;index"`
	RoleID       string `json:"role_id" gorm:"not null;index"`
	RefreshToken string `json:"refresh_token" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}

	hashedPassword, err := utils.HashPassword([]byte(u.Password))
	if err != nil {
		return errors.Wrap(err, "User.BeforeCreate")
	}
	u.Password = hashedPassword

	return nil
}
