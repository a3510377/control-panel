package database

import (
	"github.com/a3510377/control-panel/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct{ *gorm.DB }

func NewDB(filename string) (*DB, error) {
	db, err := connect()

	// Instance
	db.AutoMigrate(&models.Instance{}, &models.InstanceTags{})

	return &DB{db}, err
}

func connect() (*gorm.DB, error) {
	// TODO add else database types
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}
