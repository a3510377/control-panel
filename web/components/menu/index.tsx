import CustomLink from '../utils/CustomLink';

export interface Link {
  name: string;
  to?: string;
  links?: Link[];
}

export default function Header() {
  const links: Link[] = [
    {
      name: '基礎資料',
      links: [
        { name: '實例', to: '/instances' },
        { name: '用戶', to: '/users' },
      ],
    },
  ];

  return (
    <aside>
      <ul role="menubar">
        <div className="icon"></div>
        <li>
          <div></div>
          <ul>
            <li></li>
          </ul>
        </li>
      </ul>
    </aside>
  );
}
