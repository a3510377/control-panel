package database

import (
	"log"
	"os"

	"github.com/a3510377/control-panel/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
	Validate *validator.Validate
}

func NewDB(filename string) (*DB, error) {
	db, err := connect()

	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
		// LogLevel: logger.Warn,
		Colorful: true,
	})
	db.Config.Logger = dbLogger

	// db.Statement.Schema.LookUpField()

	// Instance
	db.AutoMigrate(&models.Instance{}, &models.Tag{})
	// Account
	db.AutoMigrate(&models.Account{})

	validate := validator.New()
	return &DB{db, validate}, err
}

func connect() (*gorm.DB, error) {
	// TODO add else database types
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

/* ----------   utils   ---------- */

func (db *DB) PreloadAll(args ...any) (tx *gorm.DB) { return db.Preload(clause.Associations, args...) }

/* ---------- utils end ---------- */

// db.Find 的快捷方法
func find[T any](db *gorm.DB, data T, id ...any) T {
	db.Find(&data, id...)
	return data
}
