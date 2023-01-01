import Layout from '#/layout';
import {
  Checkbox,
  TableCell,
  TableHead,
  TableRow,
  TableSortLabel,
  Toolbar,
  Typography,
} from '@mui/material';
import { useState } from 'react';

export default function UsersPage() {
  const [users, setUsers] = useState([]);

  const data: string[] = [];
  return (
    <Layout rootStyle={{ margin: '1em', width: '100%' }}>
      <TableHead>
        <TableRow>
          <TableCell>
            <Checkbox></Checkbox>
          </TableCell>
          {/* {data.map((d, id) => (
            <TableCell key={id}>
              <TableSortLabel></TableSortLabel>
            </TableCell>
          ))} */}
        </TableRow>
        {/* <Typography></Typography> */}
      </TableHead>
    </Layout>
  );
}
