import Link from 'next/link';
import { AutoLink } from './group';

export interface LinkItem {
  name: string;

  to: string;
}

export interface LinkGroup {
  name: string;
  links: Link[];
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

export default function Navbar() {
  return <AutoLink items={links} />;
}
