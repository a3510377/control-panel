package database

import (
	"encoding/json"
	baseErr "errors"
	"net/http"
	"strings"
	"time"

	"github.com/a3510377/control-panel/errors"
	"github.com/a3510377/control-panel/models"
	"github.com/a3510377/control-panel/service/id"
	"github.com/a3510377/control-panel/service/secret"
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

// TODO add limit
func (db *DB) GetUsers() ([]*DBAccount, error) {
	data := []models.Account{}
	users := []*DBAccount{}

	if err := db.PreloadAll().Find(&data).Error; err != nil {
		return nil, err
	}

	for _, user := range data {
		users = append(users, &DBAccount{db, user})
	}

	return users, nil
}

// 創建一個新的使用者
func (db *DB) CreateNewUser(user NewAccountData) (*DBAccount, error) {
	if err := CheckJSONData(user); err != nil {
		return nil, err
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

// state: `200` OK
// state: `401` Unauthorized
func (d *DB) GetUserByJWT(token string) (data *DBAccount, status int) {
	claims, status := secret.JWT(token).Info()
	if status != http.StatusOK {
		return nil, status
	}
	if data := d.GetUserByID(claims.ID); data != nil {
		return data, http.StatusOK
	}
	return nil, http.StatusUnauthorized
}

// 通過名稱獲取使用者
func (db *DB) GetUserByName(username string) *DBAccount {
	data := models.Account{}
	err := db.Where("name = ?", strings.ToLower(username)).First(&data).Error
	if baseErr.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &DBAccount{db, data}
}

// 通過 ID 獲取使用者
func (db *DB) GetUserByID(id id.ID) *DBAccount {
	data := &models.Account{}
	if db.First(data, id).Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &DBAccount{db, *data}
}

/* DBAccount */

func (i *DBAccount) GetNow() {
	if data := i.Db.GetUserByID(i.ID); data != nil {
		i.Account = data.Account
	}
}
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

func (d *DBAccount) CreateNewJWT() (*secret.RefreshToken, int) {
	return secret.Create(secret.Claims{
		ID:       d.ID,
		Username: d.Name,
	}, time.Hour*24*7*4) // TODO set time from config
}

func (d *DBAccount) JSON() (result map[string]any) {
	d.GetNow()
	if d.Account.Instances == nil {
		d.Account.Instances = []*models.Instance{}
	}
	b, _ := json.Marshal(&d.Account)
	_ = json.Unmarshal(b, &result)
	return
}
