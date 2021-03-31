package interfaces

import (
	"github.com/jinzhu/gorm"
)

type IDatabase interface {
	GetInstance() *gorm.DB
}
