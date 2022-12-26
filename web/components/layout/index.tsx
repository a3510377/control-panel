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
        className={style.baseNavbar}
        sx={{
          display: 'flex',
          justifyContent: 'space-between',
          height: '100vh',
        }}
      >
        <Navbar />
        {children}
      </Box>
    </>
  );
}
