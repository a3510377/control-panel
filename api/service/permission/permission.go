package permission

type PermissionType string

const (
	S PermissionType = "system"   // System Permission ( 系統 )
	I PermissionType = "instance" // Instance Permission ( 實例 )
)

const (
	ADMINISTRATOR = 1 << iota // 0 最高管理 S,I

	ManageUser // 1 管理用戶 S,I
	ManageNick // 2 管理暱稱 S,I
	ChangeNick // 3 修改自己暱稱 S,I

	ManageFile  // 4 管理實例 ( 變更檔案 ) I
	ManageEnv   // 5 管理環境變數 I
	ReadEnv     // 5 讀取環境變數 I
	ReadFile    // 6 管理實例 ( 讀取檔案 ) I
	ChangeState // 7 更改實例狀態 ( 開關機 ) I
	SendCommand // 8 發送命令 I

	ReadState // 9 讀取狀態 S,I
	ViewLog   // 10 查看日誌 S,I
)
