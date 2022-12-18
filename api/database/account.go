package database

import (
	baseErr "errors"
	"fmt"
	"strings"

	"github.com/a3510377/control-panel/errors"
	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type NewAccountData struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type DBAccount struct {
	Db *DB
	models.Account
}

// 加密密碼
func PasswordEncryption(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// 確認密碼
func HasPassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

// 創建一個新的使用者
func (db *DB) CreateNewUser(user NewAccountData) (*DBAccount, error) {
	err := db.Validate.Struct(user)
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

		return nil, baseErr.New(errorMsg)
	}

	data := models.Account{
		Name:     user.Username,
		Password: PasswordEncryption(user.Password),
	}

	if err := db.Create(&data).Error; err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return nil, errors.ErrAccountIsUse
		}
		return nil, err
	}

	return &DBAccount{db, data}, nil
}

// 通過名稱獲取使用者
func (db *DB) GetUserByName(username string) *DBAccount {
	data := &models.Account{}
	err := db.Where("name = ?", username).First(data).Error
	if baseErr.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &DBAccount{db, *data}
}

// 通過 ID 獲取使用者
func (db *DB) GetUserByID(id id.ID) *DBAccount {
	var data *models.Account
	db.First(data, id)
	if data == nil {
		return nil
	}
	return &DBAccount{db, *data}
}

/* DBAccount */

func (i *DBAccount) GetNow()                                  { i.Account = i.Db.GetUserByID(i.ID).Account }
func (i *DBAccount) Get() *gorm.DB                            { return i.Db.Model(&models.Account{ID: i.ID}) }
func (i *DBAccount) Update(column string, value any) *gorm.DB { return i.Get().Update(column, value) }

func (i *DBAccount) Updates(values any) *gorm.DB {
	return i.Get().Omit("ID").Omit("Name").Omit("CreatedAt").Updates(values)
}

func (d *DBAccount) CheckPassword(password string) bool {
	return HasPassword(password, []byte(d.Password))
}

func (d *DBAccount) UpdatePassword(password string) {
	d.Password = PasswordEncryption(password)
	d.Db.Save(&d.Account)
}
