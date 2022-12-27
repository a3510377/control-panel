import Link from 'next/link';
import { AutoLink } from './group';
import { HTMLAttributes } from 'react';

import SettingsIcon from '@mui/icons-material/Settings';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import BarChartIcon from '@mui/icons-material/BarChart';
import AppsIcon from '@mui/icons-material/Apps';
import { Box, BoxProps } from '@mui/material';

const links: Link[] = [
  {
    name: '基礎功能',
    links: [
      { name: '數據監控', to: '/', icon: <BarChartIcon /> },
      { name: '實例管理', to: '/instances/', icon: <AppsIcon /> },
      { name: '用戶管理', to: '/users/', icon: <ManageAccountsIcon /> },
    ],
  },
  {
    name: '高級功能',
    links: [{ name: '設置', to: '/settings/', icon: <SettingsIcon /> }],
  },
];

export default function Navbar(props?: HTMLAttributes<BoxProps>) {
  return (
    <Box>
      <h1 style={{ width: '100%', textAlign: 'center' }}>管理系統</h1>
      <AutoLink items={links} />
    </Box>
  );
}

export interface LinkItem {
  name: string;
  to: string;
  icon?: React.ReactElement;
}

export interface LinkGroup {
  name: string;
  links: Link[];
  icon?: React.ReactElement;
}

export type Link = LinkItem | LinkGroup;
