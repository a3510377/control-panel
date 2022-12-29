import { ReactNode } from 'react';

import Navbar from '../Navbar';
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
        <Navbar />

        {children}
      </Box>
    </>
  );
}
