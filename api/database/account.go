package database

import (
	"github.com/a3510377/control-panel/errors"
	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
func (db *DB) CreateNewUser(username string, password string) (*DBAccount, error) {
	if db.GetUserByName(username) != nil {
		return nil, errors.ErrAccountIsUse
	}

	data := models.Account{
		Name:     username,
		Password: PasswordEncryption(password),
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &DBAccount{db, data}, nil
}

// 通過名稱獲取使用者
func (db *DB) GetUserByName(username string) *DBAccount {
	var data *models.Account
	db.First(data, "name = ?", username)
	return &DBAccount{db, *data}
}

// 通過 ID 獲取使用者
func (db *DB) GetUserByID(id id.ID) *DBAccount {
	var data *models.Account
	db.First(data, "id = ?", id)
	return &DBAccount{db, *data}
}

/* DBAccount */

func (i *DBAccount) GetNow()                                  { i.Account = i.Db.GetUserByID(i.ID).Instance }
func (i *DBAccount) Get() *gorm.DB                            { return i.Db.Model(&models.Account{ID: i.ID}) }
func (i *DBAccount) Update(column string, value any) *gorm.DB { return i.Get().Update(column, value) }

func (d *DBAccount) CheckPassword(password string) bool {
	return HasPassword(password, []byte(d.Password))
}

func (d *DBAccount) UpdatePassword(password string) {
	d.Password = PasswordEncryption(password)
	d.Db.Save(&d.Account)
}
