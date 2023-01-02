import Layout from '#/layout';
import { GetUsers, User } from '@/user/user';
import {
  Checkbox,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import { useEffect, useState } from 'react';

const headCells: {
  [key: string]: { label: string };
} = {
  id: { label: 'ID' },
  name: { label: '使用者名稱' },
  nick: { label: '使用者暱稱' },
  permission: { label: '權限' },
  createTime: { label: '創建時間' },
};

const cellsValues = Object.keys(headCells) as (keyof User)[];

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    GetUsers().then((data) => setUsers(data));
  }, []);

  return (
    <Layout rootStyle={{ margin: '1em', width: '100%' }}>
      <TableContainer component={Paper}></TableContainer>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell padding="checkbox">
              <Checkbox />
            </TableCell>

            {Object.entries(headCells).map(([key, value]) => (
              <TableCell key={key} align="center">
                {value.label}
              </TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((user) => (
            <TableRow key={user.id.toString()}>
              <TableCell padding="checkbox">
                <Checkbox />
              </TableCell>

              {cellsValues.map((key) => (
                <TableCell key={key} align="center">
                  {user[key]?.toString()}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Layout>
  );
}
