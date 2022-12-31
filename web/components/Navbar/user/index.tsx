import { Avatar, Box, ButtonBase, Card } from '@mui/material';
import SettingsIcon from '@mui/icons-material/Settings';

export default function User() {
  return (
    <ButtonBase
      component="div"
      sx={{
        textAlign: 'left',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-start',
        margin: '0 1em 12px',
        minWidth: '8em',
        padding: '8px',
        borderRadius: '5px',
        cursor: 'pointer',
        '&:hover': { bgcolor: 'rgba(79,84,92, 0.6)' },
      }}
    >
      <Avatar src="" sx={{ width: '32px', height: '32px' }} />
      <Box sx={{ paddingLeft: '1em' }}>
        <span>{'test'}</span>
      </Box>
    </ButtonBase>
  );
}
