package database

import (
	"errors"
	"log"
	"os"

	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/utils/JValidator"
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

	// Instance
	db.AutoMigrate(&models.Instance{}, &models.Tag{})
	// Account
	db.AutoMigrate(&models.Account{})

	return &DB{db, JValidator.Validate}, err
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

func CheckJSONData(data any) error { return formatValidatorError(JValidator.Validate.Struct(data)) }

func formatValidatorError(err error) error {
	if err != nil && len(err.(validator.ValidationErrors)) > 0 {
		err := err.(validator.ValidationErrors)[0]

		return errors.New(err.Translate(JValidator.Trans))
	}

	return nil
}
