import Layout from '#/layout';
import { GetUsers, User } from '@/user/user';
import { Search as SearchIcon } from '@mui/icons-material';
import {
  Box,
  Checkbox,
  InputBase,
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
  create_at: { label: '創建時間' },
};

const cellsValues = Object.keys(headCells) as (keyof User)[];

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [selects, setSelects] = useState<string[]>([]);

  useEffect(() => {
    GetUsers().then((data) => setUsers(data));
  }, []);

  return (
    <Layout
      rootStyle={{
        padding: '2em',
        width: '100%',
        display: 'flex',
        justifyContent: 'center',
      }}
    >
      <TableContainer
        component={Paper}
        elevation={3}
        sx={{ maxWidth: '100em' }}
      >
        <Box>
          <Box
            sx={{
              display: 'inline-block',
              bgcolor: '#4e4e4e38',
              borderRadius: 2,
            }}
          >
            <Box
              sx={{
                display: 'flex',
                justifyContent: 'space-between',
                alignItems: 'center',
                padding: '2px 8px',
              }}
            >
              <SearchIcon sx={{ height: '100%', pointerEvents: 'none' }} />
              <InputBase placeholder="Search..." sx={{ padding: '0 8px' }} />
            </Box>
          </Box>
        </Box>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell padding="checkbox">
                <Checkbox
                  checked={users.length === selects.length}
                  indeterminate={
                    selects.length > 0 && selects.length < users.length
                  }
                  onClick={() => {
                    setSelects(
                      users.length === selects.length
                        ? []
                        : users.map((d) => d.id.toString())
                    );
                  }}
                />
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
                      onChange={(e) => {
                        const id = user.id.toString();

                        setSelects(
                          e.target.checked
                            ? selects.concat(id)
                            : selects.filter((item) => item !== id)
                        );
                      }}
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
    </Layout>
  );
}
