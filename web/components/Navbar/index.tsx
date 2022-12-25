import Link from 'next/link';
import { ListLinkItem } from './item';
export interface Link {
  name: string;
  to?: string;

  links?: Link[];
}
const links: Link[] = [
  {
    name: 'Home',
    links: [
      { name: 'Home', to: '/' },
      { name: 'About', to: '/about' },
    ],
  },
];

export default function Navbar() {
  return (
    <>
      <ListLinkItem text="awa" to="/a" />
    </>
  );
}
