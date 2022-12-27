import { Box, Button, Link, TextField } from '@mui/material';

export default function RegisterPage() {
  return (
    <Box component="form" noValidate sx={{ mt: 3 }}>
      <TextField
        label="綽號"
        autoComplete="nickname"
        name="nickname"
        id="nickname"
        required
        fullWidth
        autoFocus
      />
      <TextField
        label="帳號"
        autoComplete="email"
        name="email"
        id="email"
        required
        fullWidth
      />
      <TextField
        label="密碼"
        name="password"
        type="password"
        id="password"
        autoComplete="new-password"
        required
        fullWidth
      />
      <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
        註冊
      </Button>
      <Link href="#" variant="body2">
        已有帳戶? 登入
      </Link>
    </Box>
  );
}
