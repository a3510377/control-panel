import { ReactNode } from 'react';

import Navbar from '../Navbar';
import {
  Box,
  SxProps,
  Theme,
  ThemeOptions,
  ThemeProvider,
  createTheme,
} from '@mui/material';

export interface Props {
  children?: ReactNode;
  theme?: ThemeOptions;
  themeData?: Theme;
  rootStyle?: SxProps<Theme>;
}

export default function Layout({
  children,
  theme,
  themeData,
  rootStyle,
}: Props) {
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

        <Box
          className="scrollbar-style"
          sx={{ overflow: 'auto', ...rootStyle }}
        >
          {children}
        </Box>
      </Box>
    </ThemeProvider>
  );
}
