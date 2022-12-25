import {
  Link,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from '@mui/material';

export interface ListLinkItemProps {
  text: string;
  to: string;
  icon?: React.ReactElement;
  index?: number;
}

export function ListLinkItem({ icon, text, to, index }: ListLinkItemProps) {
  return (
    <ListItemButton
      LinkComponent={Link}
      href={to}
      style={{ borderRadius: '5px' }}
      sx={{ pl: index ? index * 4 : void 0 }}
    >
      {icon && <ListItemIcon>{icon}</ListItemIcon>}
      <ListItemText primary={text} />
    </ListItemButton>
  );
}
