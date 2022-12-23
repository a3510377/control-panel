import CustomLink from '../utils/CustomLink';

export default function Header() {
  return (
    <header>
      <nav>
        <div>
          <CustomLink href="">Home</CustomLink>
        </div>
        <div className="">
          <CustomLink href="/login">Logout</CustomLink>
        </div>
      </nav>
    </header>
  );
}
