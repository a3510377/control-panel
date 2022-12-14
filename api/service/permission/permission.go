package permission

type (
	PermissionType string
	Permission     int
)

const (
	S PermissionType = "system"   // System Permission ( 系統 )
	I PermissionType = "instance" // Instance Permission ( 實例 )
)

const (
	Administrator = 1 << iota // 1 最高管理 S,I

	ManageUser // 2 管理用戶 S,I
	ManageNick // 3 管理暱稱 S,I
	ChangeNick // 4 修改自己暱稱 S,I

	ManageFile  // 5 管理實例 ( 變更檔案 ) I
	ManageEnv   // 6 管理環境變數 I
	ReadEnv     // 7 讀取環境變數 I
	ReadFile    // 8 管理實例 ( 讀取檔案 ) I
	ChangeState // 9 更改實例狀態 ( 開關機 ) I
	SendCommand // 10 發送命令 I

	ReadState // 11 讀取狀態 S,I
	ViewLog   // 12 查看日誌 S,I

	None = 0 // 0 無權限 S,I
)

// if self has the same or fewer permissions as other.
func (p Permission) Subset(other Permission) bool {
	return (p & other) == p
}

// if self has the same or more permissions as other.
func (p Permission) Superset(other Permission) bool {
	return (p | other) == p
}

// if the permissions on other are a strict subset of those on self.
func (p Permission) StrictSubset(other Permission) bool {
	return p.Subset(other) && p != other
}

// if the permissions on other are a strict superset of those on self.
func (p Permission) StrictSuperset(other Permission) bool {
	return p.Superset(other) && p != other
}
