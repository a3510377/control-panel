import {
  Link,
  ListItemButton,
  ListItemIcon,
  ListItemText,
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
    <ListItemButton
      selected={router.asPath === to}
      LinkComponent={Link}
      href={to}
      sx={{
        borderRadius: '5px',
        mb: '5px',
        pl: index ? index * 4 : void 0,
        '&:hover': { bgcolor: '#eef5fb17' },
      }}
    >
      {icon && <ListItemIcon>{icon}</ListItemIcon>}
      <ListItemText primary={text} />
    </ListItemButton>
  );
}
