import Head from 'next/head';
import { ChangeEvent, FormEventHandler, useState } from 'react';

import style from './index.module.scss';
import { Login, LoginErrorType, LoginInfo } from '../../api/user';
import classNames from 'classnames';
import Router from 'next/router';

export default function Home() {
  const [hasInputName, setHasInputName] = useState('');
  const [hasInputPassword, setHasInputPassword] = useState('');
  const [checkInputError, setCheckInputError] = useState('');
  const [errMessage, setErrMessage] = useState('');

  const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    const loginData = await Login(hasInputName, hasInputPassword);

    if (loginData?.type === 'success') {
      localStorage.setItem('token', (loginData as LoginInfo).token.token);
      Router.push('/');
    } else {
      const data = loginData as LoginErrorType;

      setErrMessage(data.error);
      setCheckInputError(data.type);
    }
  };

  const updateInput = (
    callable: (value: string) => void,
    e: ChangeEvent<HTMLInputElement>
  ) => {
    setCheckInputError('');
    callable(e.target.value);
  };

  return (
    <>
      <Head>
        <title>登入 - 管理系統</title>
      </Head>
      <div className={style.main}>
        <div>
          <h1>登入</h1>
          <form onSubmit={handleSubmit}>
            <div className={style.inputBox}>
              {checkInputError === 'username' && (
                <div className="error-message">{errMessage}</div>
              )}

              <input
                className={classNames(
                  hasInputName && 'input',
                  checkInputError === 'username' && 'error'
                )}
                onChange={updateInput.bind(null, setHasInputName)}
                aria-label="用戶名"
                type="text"
                id="login_field"
                autoCorrect="off"
                autoCapitalize="off"
                autoComplete="username"
                autoFocus
              />
              <label htmlFor="login_field">用戶名</label>
              <div className="select-box" />
            </div>
            <div className={style.inputBox}>
              {checkInputError === 'password' && (
                <div className="error-message">{errMessage}</div>
              )}

              <input
                className={classNames(
                  hasInputPassword && 'input',
                  checkInputError === 'password' && 'error'
                )}
                onChange={updateInput.bind(null, setHasInputPassword)}
                aria-label="密碼"
                type="password"
                name="password"
                id="password"
                autoComplete="current-password"
              />
              <label htmlFor="password">密碼</label>
              <div className="select-box" />
            </div>
            <div className={style.loginBox}>
              <a href="#" className="forgot">
                忘記密碼
              </a>

              <button disabled={!hasInputName || !hasInputPassword}>
                登入
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  );
}
