package models

// Role model role
type Role struct {
	Model       `json:"inline"`
	Name        string `json:"name" gorm:"unique;not null;index"`
	Description string `json:"description"`
}
