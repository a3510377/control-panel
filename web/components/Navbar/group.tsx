import {
  Box,
  Collapse,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from '@mui/material';
import { useState } from 'react';
import { ExpandLess, ExpandMore } from '@mui/icons-material';

import { ListLinkItem } from './item';
import { Link } from '.';

export function AutoLink({
  items,
  index: LIndex,
}: {
  items: Link[];
  index?: number;
}) {
  return (
    <>
      {items?.map((item, index) =>
        'to' in item ? (
          <ListLinkItem
            key={index}
            text={item.name}
            to={item.to}
            index={LIndex}
            icon={item.icon}
          />
        ) : (
          <ListGroup
            key={index}
            title={item.name}
            items={item.links}
            index={LIndex}
            icon={item.icon}
          />
        )
      )}
    </>
  );
}

export interface ListGroupProps {
  title: string;
  items: Link[];
  index?: number;
  icon?: React.ReactElement;
}

export default function ListGroup({
  title,
  items,
  index,
  icon,
}: ListGroupProps) {
  const [open, setOpen] = useState(true);

  return (
    <>
      <ListItemButton
        onClick={setOpen.bind(null, !open)}
        sx={{ pl: index ? index * 4 : void 0 }}
      >
        {icon && <ListItemIcon>{icon}</ListItemIcon>}

        <ListItemText primary={title} />
        {open ? <ExpandLess /> : <ExpandMore />}
      </ListItemButton>

      <Collapse in={open} timeout="auto" unmountOnExit>
        {items && <AutoLink items={items} index={(index || 0) + 1} />}
      </Collapse>
    </>
  );
}
