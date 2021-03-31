package dbs

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/config"
)

type database struct {
	db *gorm.DB
}

func NewDatabase() interfaces.IDatabase {
	dbConfig := config.Config.Database
	connectionPath := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password, dbConfig.SSLMode)

	logger.Info(connectionPath)

	db, err := gorm.Open("postgres", connectionPath)
	if err != nil {
		logger.Fatal("Cannot connect to database: ", err)
	}

	// Set up connection pool
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)

	return &database{
		db: db,
	}
}

func (d *database) GetInstance() *gorm.DB {
	return d.db
}
