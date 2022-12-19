package database

import (
	baseErr "errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/a3510377/control-panel/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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

func CheckJSONData(user any) error {
	err := validate.Struct(user)
	if err != nil && len(err.(validator.ValidationErrors)) > 0 {
		err := err.(validator.ValidationErrors)[0]
		errorMsg := strings.ToLower(err.Field())

		switch err.Tag() {
		case "required":
			errorMsg += " is required."
		case "min":
			errorMsg += fmt.Sprintf(" (%v) required at least %v", err.Type(), err.Param())
		case "max":
			errorMsg += fmt.Sprintf(" (%v) only has a maximum of %v", err.Type(), err.Param())
		default:
			errorMsg = err.Error()
		}

		return baseErr.New(errorMsg)
	}

	return nil
}
