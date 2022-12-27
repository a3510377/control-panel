import Layout from '#/layout';
import { useRouter } from 'next/router';

export default function InstancePage() {
  const router = useRouter();
  const { id } = router.query;

  return (
    <Layout>
      <div>{id}</div>
    </Layout>
  );
}
