import Head from 'next/head';
import { ChangeEvent, FormEventHandler, useRef, useState } from 'react';
import Router from 'next/router';
import { TextField } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import LoginIcon from '@mui/icons-material/Login';

import style from './index.module.scss';
import { Login, LoginErrorType, LoginInfo } from '@/user';

export default function LoginPage() {
  const formRef = useRef<HTMLFormElement>(null);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [errMessage, setErrMessage] = useState('');
  const [errType, setErrType] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    setLoading(true);
    const loginData = await Login(username, password);

    if (loginData?.type === 'success') {
      localStorage.setItem('token', (loginData as LoginInfo).token.token);
      return Router.push('/');
    } else {
      const data = loginData as LoginErrorType;

      setErrMessage(data.error);
      setErrType(data.type);
    }
    setLoading(false);
  };

  const updateInputData = (
    callback: (_: string) => void,
    e: ChangeEvent<HTMLInputElement>
  ) => {
    callback(e.target.value);
    errType === e.target.id && setErrType('');
  };

  return (
    <>
      <Head>
        <title>登入 - 管理系統</title>
        <meta name="description" content={void 0} />
      </Head>
      <div className={style.main}>
        <form onSubmit={handleSubmit} ref={formRef}>
          <h1>登入</h1>
          <TextField
            error={errType === 'username'}
            helperText={errType === 'username' && errMessage}
            margin="normal"
            required
            fullWidth
            id="email"
            label="帳號"
            autoComplete="email"
            autoFocus
            value={username}
            onChange={updateInputData.bind(null, setUsername)}
            inputProps={{ tabIndex: 1 }}
          />
          <TextField
            error={errType === 'password'}
            helperText={errType === 'password' && errMessage}
            margin="normal"
            required
            fullWidth
            name="password"
            label="密碼"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={updateInputData.bind(null, setPassword)}
            inputProps={{ tabIndex: 2 }}
          />
          <div className={style.loginBox}>
            <a href="#" className="forgot" tabIndex={4}>
              忘記密碼
            </a>
            <LoadingButton
              type="submit"
              loading={loading}
              variant="outlined"
              startIcon={<LoginIcon />}
              tabIndex={3}
            >
              登入
            </LoadingButton>
          </div>
        </form>
      </div>
    </>
  );
}
