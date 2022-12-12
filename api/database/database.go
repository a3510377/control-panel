package database

import (
	"log"
	"os"

	"github.com/a3510377/control-panel/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct{ *gorm.DB }

func NewDB(filename string) (*DB, error) {
	db, err := connect()

	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Warn,
		Colorful: false,
	})
	db.Config.Logger = dbLogger

	// Instance
	db.AutoMigrate(&models.Instance{}, &models.InstanceTags{})

	return &DB{db}, err
}

func connect() (*gorm.DB, error) {
	// TODO add else database types
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}
