package models

import (
	"crypto/md5"
	"fmt"

	"github.com/a3510377/control-panel/service/id"
	"gorm.io/gorm"
)

func MD5(str string) string {
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has) // hex
}

type Account struct {
	ID         id.ID       `json:"id" gorm:"primarykey"`         // ID
	Name       string      `json:"name" gorm:"size:60;not null"` // 實例名稱
	Nick       string      `json:"nick" gorm:"size:60;not null"` // 暱稱
	Password   string      `json:"-" gorm:"size:60;not null"`    // 密碼
	Permission string      `json:"permission"`                   // 權限
	Instances  []*Instance `gorm:"many2many"`

	// TODO Email    string `json:"email" gorm:"size:60;not null"` // 電子郵件
}

func NewAccount() *Instance {
	return &Instance{}
}

func (i *Account) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = id.GlobalIDMake.Generate()
	return
}
