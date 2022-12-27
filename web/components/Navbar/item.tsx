import {
  Link,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  ThemeProvider,
  createTheme,
} from '@mui/material';
import { useRouter } from 'next/router';

export interface ListLinkItemProps {
  text: string;
  to: string;
  icon?: React.ReactElement;
  index?: number;
}

export function ListLinkItem({ icon, text, to, index }: ListLinkItemProps) {
  const router = useRouter();

  return (
    <ThemeProvider
      theme={createTheme({
        components: {
          MuiListItemButton: {
            styleOverrides: {
              root: {
                borderRadius: '5px',
                marginBottom: '8px',
                ':hover': { backgroundColor: '#383838ad' },
                '&.Mui-selected': {
                  backgroundColor: '#363d4ead',
                  ':hover': { backgroundColor: '#2b3140d4' },
                },
              },
            },
          },
        },
      })}
    >
      <ListItemButton
        selected={router.asPath === to}
        LinkComponent={Link}
        href={to}
        sx={{ pl: index ? index * 4 : void 0 }}
      >
        {icon && <ListItemIcon>{icon}</ListItemIcon>}
        <ListItemText primary={text} />
      </ListItemButton>
    </ThemeProvider>
  );
}
