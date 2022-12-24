import Head from 'next/head';

import './globals.scss';

import type { AppProps } from 'next/app';
import { Inter } from '@next/font/google';

import { useEffect } from 'react';
import Authorization from '../components/utils/Authorization';

const inter = Inter({ subsets: ['latin'] });

export default function MyApp({ Component, pageProps }: AppProps) {
  useEffect(() => {
    const loop = handleDevTools();

    return () => clearInterval(loop);
  }, []);

  return (
    <>
      <Head>
        <title>管理系統</title>
      </Head>
      <Authorization />
      <main className={inter.className}>
        <Component {...pageProps} />
      </main>
    </>
  );
}

const handleDevTools = (threshold = 160) => {
  const devtools: {
    isOpen: boolean;
    orientation?: 'vertical' | 'horizontal';
  } = { isOpen: false, orientation: undefined };

  const emitEvent = (isOpen: boolean, orientation?: string) => {
    dispatchEvent(
      new CustomEvent('devtools-change', { detail: { isOpen, orientation } })
    );
  };

  const loop = () => {
    const widthThreshold = outerWidth - innerWidth > threshold;
    const heightThreshold = outerHeight - innerHeight > threshold;
    const orientation = widthThreshold ? 'vertical' : 'horizontal';

    if (
      !(heightThreshold && widthThreshold) &&
      (widthThreshold || heightThreshold)
    ) {
      if (!devtools.isOpen || devtools.orientation !== orientation)
        emitEvent(true, orientation);

      devtools.isOpen = true;
      devtools.orientation = orientation;
    } else {
      if (devtools.isOpen) emitEvent(false, undefined);

      devtools.isOpen = false;
      devtools.orientation = undefined;
    }
  };

  return setInterval(loop, threshold);
};
