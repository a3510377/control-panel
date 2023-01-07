import { Button, Box } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';
import AddIcon from '@mui/icons-material/Add';
import Grid from '@mui/material/Unstable_Grid2';

export enum Permission {}
// // 1 最高管理 S,I
// Administrator = 1 << 0,
// // 2 管理用戶 S,I
// ManageUser = 1 << 0,
// // 3 管理暱稱 S,I
// ManageNick = 1 << 0,
// // 4 修改自己暱稱 S,I
// ChangeNick = 1 << 0,
// // 5 管理實例 ( 變更檔案 ) I
// ManageFile = 1 << 0,
// // 6 管理環境變數 I
// ManageEnv = 1 << 0,
// // 7 讀取環境變數 I
// ReadEnv = 1 << 0,
// // 8 管理實例 ( 讀取檔案 ) I
// ReadFile = 1 << 0,
// // 9 更改實例狀態 ( 開關機 ) I
// ChangeState = 1 << 0,
// // 10 發送命令 I
// SendCommand = 1 << 0,
// // 11 讀取狀態 S,I
// ReadState = 1 << 0,
// // 12 查看日誌 S,I
// ViewLog = 1 << 0,
// // 0 無權限 S,I
// None = 0,

export const permissions = [];

export default function Operate() {
  return (
    <Grid container gap={1}>
      <CreateUser />
      <DeleteUser />
    </Grid>
  );
}

export function DeleteUser() {
  return (
    <Button variant="contained" color="error" startIcon={<DeleteIcon />}>
      刪除用戶
    </Button>
  );
}

export function CreateUser() {
  return (
    <>
      <Button variant="contained" color="success" startIcon={<AddIcon />}>
        添加用戶
      </Button>
      <Box></Box>
    </>
  );
}

export function PermissionUser() {
  return;
}

// export function () {
//   return;
// }
