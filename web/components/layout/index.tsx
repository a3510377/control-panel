import { ReactNode } from 'react';

import Navbar from '../Navbar';

export interface Props {
  children?: ReactNode;
}

export default function Layout({ children }: Props) {
  return (
    <>
      <Navbar />
      {children}
    </>
  );
}
