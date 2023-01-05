import Layout from '#/layout';
import { GetUsers, User } from '@/user/user';
import {
  Box,
  Checkbox,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
} from '@mui/material';
import AutorenewIcon from '@mui/icons-material/Autorenew';
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

  const getUsers = () => GetUsers().then((data) => setUsers(data));
  useEffect(() => {
    getUsers();
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
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            padding: '0 1em',
          }}
        >
          <TextField
            helperText={''}
            margin="normal"
            fullWidth
            name="search"
            label="搜尋"
            id="search"
            // onChange={updateInputData.bind(null, setPassword)}
            inputProps={{ tabIndex: 1 }}
            type="search"
            sx={{ width: '240px' }}
          />
          <IconButton sx={{ marginLeft: '10px' }} onClick={getUsers}>
            <AutorenewIcon />
          </IconButton>
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
                <TableCell key={`table-table-${key}`} align="center">
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
                  key={`table-body-${user.id.toString()}`}
                  role="checkbox"
                  selected={isSelect}
                  tabIndex={-1}
                  onClick={(e) => {
                    const id = user.id.toString();

                    if (e.ctrlKey) {
                      e.preventDefault();

                      setSelects(
                        selects.includes(id)
                          ? selects.filter((item) => item !== id)
                          : selects.concat(id)
                      );
                    }
                  }}
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
                    <TableCell key={`body-table-cell-${key}`} align="center">
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
