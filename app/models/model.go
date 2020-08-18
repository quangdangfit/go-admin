package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        string    `json:"id" gorm:"unique;not null;index;primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;index"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;index"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	model.ID = uuid.New().String()
	model.CreatedAt = time.Now().UTC()
	model.UpdatedAt = time.Now().UTC()
	return nil
}
