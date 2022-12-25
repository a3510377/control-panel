import {
  Link,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from '@mui/material';

export interface ListLinkItemProps {
  icon?: React.ReactElement;
  text: string;
  to: string;
}

export function ListLinkItem({ icon, text, to }: ListLinkItemProps) {
  return (
    <ListItemButton
      LinkComponent={Link}
      href={to}
      style={{ borderRadius: '5px' }}
    >
      {icon && <ListItemIcon>{icon}</ListItemIcon>}
      <ListItemText primary={text} />
    </ListItemButton>
  );
}
