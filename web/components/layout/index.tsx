import { ReactNode } from 'react';

import Navbar from '../Navbar';

export interface Props {
  children?: ReactNode;
}

export default function Layout({ children }: Props) {
  return (
    <>
      <Navbar
        style={{
          width: '240px',
          backgroundColor: '#383838',
          color: 'white',
          height: '100vh',
        }}
      />
      {children}
    </>
  );
}
