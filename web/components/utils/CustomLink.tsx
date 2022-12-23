import { useRouter } from 'next/router';
import { useEffect, ReactNode, HTMLAttributes } from 'react';

type CustomLinkProps = {
  children: ReactNode;
  href: string;
  prefetch?: boolean;
  replace?: boolean;
  shallow?: boolean;
} & HTMLAttributes<HTMLAnchorElement>;

export default function CustomLink({
  children,
  href,
  prefetch = false,
  replace = false,
  shallow = false,
  ...props
}: CustomLinkProps) {
  const router = useRouter();

  useEffect(() => {
    if (prefetch) router.prefetch(href);
  }, [router, href, prefetch]);

  return (
    <a
      {...props}
      href={href}
      onClick={(event) => {
        event.preventDefault();
        (replace ? router.replace : router.push)(href, undefined, { shallow });
      }}
    >
      {children}
    </a>
  );
}
