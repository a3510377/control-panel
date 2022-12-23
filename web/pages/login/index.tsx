import Head from 'next/head';
import { useState } from 'react';

import style from './index.module.scss';

export default function Home() {
  const [hasInputName, setHasInputName] = useState(false);
  const [hasInputPassword, setHasInputPassword] = useState(false);

  return (
    <>
      <Head>
        <title>Login</title>
        <meta name="description" content="" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={style.main}>
        <div>
          <h1>身份認證</h1>
          <form className={style.form} action="#" method="POST">
            <div className={style.inputBox}>
              <input
                className={hasInputName ? 'input' : void 0}
                onChange={(e) => setHasInputName(!!e.target.value)}
                aria-label="用戶名"
                type="text"
                id="login_field"
                autoCorrect="off"
                autoCapitalize="off"
                autoComplete="username"
                autoFocus
              />
              <label htmlFor="login_field">用戶名</label>
            </div>
            <div className={style.inputBox}>
              <input
                className={hasInputPassword ? 'input' : void 0}
                onChange={(e) => setHasInputPassword(!!e.target.value)}
                aria-label="密碼"
                type="password"
                name="password"
                id="password"
                autoComplete="current-password"
              />
              <label htmlFor="password">密碼</label>
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
      </main>
    </>
  );
}
