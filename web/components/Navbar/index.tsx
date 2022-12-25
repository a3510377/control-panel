import Link from 'next/link';
import { AutoLink } from './group';
import { HTMLAttributes } from 'react';

export interface LinkItem {
  name: string;
  to: string;
  icon?: React.ReactElement;
}

export interface LinkGroup {
  name: string;
  links: Link[];
  icon?: React.ReactElement;
}

export type Link = LinkItem | LinkGroup;

const links: Link[] = [
  {
    name: 'a',
    links: [
      { name: 'Home', to: '/' },
      { name: 'About', to: '/about' },
    ],
  },
  {
    name: 'b',
    links: [{ name: 'Home', to: '/' }],
  },
];

export default function Navbar(props?: HTMLAttributes<HTMLDivElement>) {
  return (
    <div {...props}>
      <AutoLink items={links} />
    </div>
  );
}
