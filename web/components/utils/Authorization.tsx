import Router from 'next/router';
import { useEffect } from 'react';
import { GetInfo } from '../../api/user';

const toLogin = () => Router.replace('/login', void 0);

export default function Authorization() {
  let one = 0;

  useEffect(() => {
    if (
      one++ || // Prevent multiple calls
      Router.route === '/_error' // Prevent redirect to login page when 404
    ) {
      return;
    }

    // Don't redirect to login page if already on login page
    if (!localStorage.getItem('token') && Router.asPath !== '/login/') {
      toLogin();
      return;
    }

    // Redirect to home page if already logged in
    if (localStorage.getItem('token')) {
      GetInfo()
        ?.then(() => Router.asPath === '/login/' && Router.replace('/'))
        .catch(toLogin);
    }
  }, [one]);

  return <></>;
}
