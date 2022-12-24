import Router from 'next/router';
import { useEffect } from 'react';
import { GetInfo } from '../../api/user';

const toLogin = () => Router.push('/login', void 0);

export default function Authorization() {
  let one = 0;

  useEffect(() => {
    console.log('--', one);

    if (one++) return; // Prevent multiple calls
    console.log(one);

    // Don't redirect to login page if already on login page
    if (!localStorage.getItem('token') && Router.asPath !== '/login/') {
      toLogin();
      return;
    }

    if (localStorage.getItem('token')) {
      GetInfo()
        ?.then(() => Router.asPath === '/login/' && Router.push('/'))
        .catch(toLogin);
    }
  }, [one]);

  return <></>;
}
