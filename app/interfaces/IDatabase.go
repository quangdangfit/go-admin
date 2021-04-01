package interfaces

import (
	"github.com/jinzhu/gorm"
)

// IDatabase interface
type IDatabase interface {
	GetInstance() *gorm.DB
}
