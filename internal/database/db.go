package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/JasperRosales/todo-api/internal/models"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}
	})

	return dbInstance
}

func InitDB() error {
	db := GetDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}
	return nil
}
