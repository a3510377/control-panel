import { BaseOverview, OverviewBaseData } from '@/system/overview';
import { useEffect, useState } from 'react';

export default function Info() {
  const [overviewData, setOverviewData] = useState<OverviewBaseData>();
  // const data = BaseOverview()

  useEffect(() => {
    BaseOverview()?.then((data) => {
      if (data && !('error' in data)) setOverviewData(data);
    });
  }, []);
  return (
    <div>
      <h1>伺服器詳細數據</h1>
      <div></div>
    </div>
  );
}
