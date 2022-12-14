package models

import (
	"github.com/a3510377/control-panel/service/id"
	"github.com/a3510377/control-panel/utils/JTime"
	"gorm.io/gorm"
)

type Instance struct {
	ID           id.ID      `json:"id" gorm:"primarykey"`                                          // ID
	Name         string     `json:"name" gorm:"size:20;not null" validate:"required,min=1,max=20"` // 實例名稱
	StartCommand string     `json:"start_command"`                                                 // 開機指令
	StopCommand  string     `json:"stop_command"`                                                  // 停止指令
	RootDir      string     `json:"root_dir"`                                                      // 根目錄
	Type         string     `json:"type"`                                                          // 實例類型
	CreatedAt    JTime.Time `json:"create_at"`                                                     // 創建時間
	LastTime     JTime.Time `json:"last_time,omitempty" gorm:"default:null"`                       // 最後一次啟動時間
	EndAt        JTime.Time `json:"end_time,omitempty" gorm:"default:null"`                        // 到期時間
	Tags         []*Tag     `json:"tags" gorm:"many2many:instanceTags" validate:"max=10"`          // 標籤
	AutoStart    bool       `json:"auto_start"`                                                    // 自動啟動
	AutoRestart  bool       `json:"auto_restart"`                                                  // 自動重啟

	// TODO add MAXRam, MAXCpu, FRP server config
}

type Tag struct {
	ID        int         `json:"-" gorm:"primarykey"`
	Name      string      `json:"name" gorm:"size:15;not null"`
	Instances []*Instance `json:"instances" gorm:"many2many:instanceTags"`
}

func NewInstance() *Instance {
	return &Instance{}
}

func (i *Instance) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = id.GlobalIDMake.Generate()
	return
}
