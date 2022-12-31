import Typography from '@mui/material/Typography';
import { ResponsiveStyleValue } from '@mui/system';
import { ReactNode } from 'react';

export interface TitleProps {
  children?: ReactNode;
  color?: ResponsiveStyleValue<string[]>;
}

export default function Title(props: TitleProps) {
  return (
    <Typography
      component="h2"
      variant="h6"
      color={props.color || 'primary'}
      gutterBottom
    >
      {props.children}
    </Typography>
  );
}
