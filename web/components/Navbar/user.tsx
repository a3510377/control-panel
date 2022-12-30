import { Avatar, Box } from '@mui/material';
import SettingsIcon from '@mui/icons-material/Settings';

export default function User() {
  return (
    <Box
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '0 20px 15px',
      }}
    >
      <Box sx={{ display: 'flex', alignItems: 'center' }}>
        <Avatar src="" />
        <Box sx={{ paddingLeft: '8px' }}>
          <span>{'test'}</span>
        </Box>
      </Box>

      <SettingsIcon
        sx={{
          color: '#dcddde',
          cursor: 'pointer',
          padding: '6px',
          width: '20px',
          height: '20px',
          borderRadius: '4px',
          '&:hover': { bgcolor: 'rgba(79,84,92, 0.6)' },
        }}
      />
    </Box>
  );
}
