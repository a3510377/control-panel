import { Avatar, Box } from '@mui/material';
import SettingsIcon from '@mui/icons-material/Settings';

export default function User() {
  return (
    <Box
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        margin: '0 1em 12px',
      }}
    >
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
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
