package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func New(filename string) (*DB, error) {
	Db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	return &DB{
		Db,
	}, err
}
