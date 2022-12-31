import { ReactNode } from 'react';

import Navbar from '../Navbar';
import {
  Box,
  Theme,
  ThemeOptions,
  ThemeProvider,
  createTheme,
} from '@mui/material';

export interface Props {
  children?: ReactNode;
  theme?: ThemeOptions;
  themeData?: Theme;
}

export default function Layout({ children, theme, themeData }: Props) {
  return (
    <ThemeProvider theme={themeData || createTheme(theme)}>
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
    </ThemeProvider>
  );
}
