package models

import (
	"strings"

	"github.com/a3510377/control-panel/service/id"
	"github.com/a3510377/control-panel/service/permission"
	"github.com/a3510377/control-panel/utils/JTime"
	"gorm.io/gorm"
)

type Account struct {
	ID         id.ID                 `json:"id" gorm:"primarykey"`                       // ID
	Name       string                `json:"name" gorm:"uniqueIndex;size:20;not null"`   // 實例名稱
	Nick       string                `json:"nick" gorm:"size:20;default:null"`           // 暱稱
	Password   string                `json:"-" gorm:"not null"`                          // 密碼
	Permission permission.Permission `json:"permission"`                                 // 權限
	CreatedAt  JTime.Time            `json:"create_at"`                                  // 創建時間
	Instances  []*Instance           `json:"instances" gorm:"many2many:instanceAccount"` //

	// TODO Email    string `json:"email" gorm:"size:60;not null"` // 電子郵件
}

func NewAccount() *Account { return &Account{} }

func (i *Account) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = id.GlobalIDMake.Generate()

	return
}

func (i *Account) BeforeSave(tx *gorm.DB) (err error) {
	i.Name = strings.ToLower(i.Name)
	// tx.Statement.SetColumn("Name", i.Name)

	return
}
