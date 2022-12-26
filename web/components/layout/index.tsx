import { ReactNode } from 'react';

import Navbar from '../Navbar';
import style from './base.module.scss';
import { Box } from '@mui/material';

export interface Props {
  children?: ReactNode;
}

export default function Layout({ children }: Props) {
  return (
    <>
      <Box
        component="main"
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          height: '100vh',
        }}
      >
        <Box className={style.baseNavbar}>
          <Navbar />
        </Box>
        {children}
      </Box>
    </>
  );
}
