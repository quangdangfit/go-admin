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
	if model.ID == "" {
		model.ID = uuid.New().String()
	}
	if model.CreatedAt.IsZero() {
		model.CreatedAt = time.Now()
	}
	if model.UpdatedAt.IsZero() {
		model.UpdatedAt = time.Now()
	}
	return nil
}

func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	model.UpdatedAt = time.Now().UTC()
	return nil
}
