package models

import (
	"strings"
	"time"

	"github.com/a3510377/control-panel/service/id"
	"github.com/a3510377/control-panel/service/permission"
	"gorm.io/gorm"
)

type Account struct {
	ID         id.ID                 `json:"id" gorm:"primarykey"`         // ID
	Name       string                `json:"name" gorm:"size:20;not null"` // 實例名稱
	Nick       string                `json:"nick" gorm:"size:20;not null"` // 暱稱
	Password   string                `json:"-" gorm:"not null"`            // 密碼
	Permission permission.Permission `json:"permission"`                   // 權限
	CreatedAt  time.Time             `json:"create_at"`                    // 創建時間
	Instances  []*Instance           `gorm:"many2many:instanceAccount"`    //

	// TODO Email    string `json:"email" gorm:"size:60;not null"` // 電子郵件
}

func NewAccount() *Account { return &Account{} }

func (i *Account) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = id.GlobalIDMake.Generate()
	i.Name = strings.ToLower(i.Name)
	return
}
