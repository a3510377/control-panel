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
  TablePagination,
  TableRow,
} from '@mui/material';
import { useEffect, useState } from 'react';

const headCells: {
  [key: string]: { label: string };
} = {
  name: { label: '使用者名稱' },
  nick: { label: '使用者暱稱' },
  id: { label: 'ID' },
  permission: { label: '權限' },
  createTime: { label: '創建時間' },
};

const cellsValues = Object.keys(headCells) as (keyof User)[];

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [selects, setSelects] = useState<string[]>([]);

  const handleClick = (id: string) => {
    return setSelects.bind(
      null,
      selects.includes(id)
        ? selects.filter((id) => id !== id)
        : [...selects, id]
    );
  };

  useEffect(() => {
    GetUsers().then((data) => setUsers(data));
  }, []);

  return (
    <Layout rootStyle={{ margin: '1em', width: '100%' }}>
      <TableContainer component={Paper}>
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
            {users.map((user) => {
              const isSelect = selects.includes(user.id.toString());

              return (
                <TableRow
                  hover
                  key={user.id.toString()}
                  role="checkbox"
                  selected={isSelect}
                  tabIndex={-1}
                >
                  <TableCell padding="checkbox">
                    <Checkbox
                      checked={isSelect}
                      onClick={handleClick(user.id.toString())}
                    />
                  </TableCell>

                  {cellsValues.map((key) => (
                    <TableCell key={key} align="center">
                      {user[key]?.toString()}
                    </TableCell>
                  ))}
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
      {/* <TablePagination
        page={0}
        count={10}
        onPageChange={(...data) => {
          return;
        }}
        rowsPerPage={10}
        // rowsPerPageOptions={[5, 10, 25]}
        // component="div"
        // count={10}
        // rowsPerPage={rowsPerPage}
        // page={page}
        // onPageChange={handleChangePage}
        // onRowsPerPageChange={handleChangeRowsPerPage}
      /> */}
    </Layout>
  );
}
