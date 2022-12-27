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
        <Box
          sx={{
            bgcolor: 'rgb(30, 30, 30)',
            width: '280px',
            zIndex: 100,
            color: 'white',
          }}
        >
          <Navbar />
        </Box>
        {children}
      </Box>
    </>
  );
}
