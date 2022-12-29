import { Avatar, Box, ListItemText } from '@mui/material';
import SettingsIcon from '@mui/icons-material/Settings';

export default function User() {
  return (
    <Box
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '0 10px 10px',
      }}
    >
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <Avatar src="" />
        <Box sx={{ 'padding-left': '8px' }}>
          <span>{'test'}</span>
        </Box>
      </Box>

      <SettingsIcon />
    </Box>
  );
}
