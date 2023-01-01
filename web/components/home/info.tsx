import {
  BaseOverview,
  OverviewBaseData,
  OverviewData,
} from '@/system/overview';
import Grid from '@mui/material/Unstable_Grid2';
import { useEffect, useState } from 'react';
import Title from './title';

export default function Info({ data }: { data: (OverviewData | void)[] }) {
  const [BaseOverviewData, setOverviewData] = useState<OverviewBaseData>();
  const [dataInfo, setDataInfo] = useState<{ label: string; data: string }[]>();

  const timeStr = (value?: string | number | Date) => {
    const date = value ? new Date(value) : new Date();
    return [
      // year month day
      [date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate()].join('/'),
      [
        date.getHours().toString().padStart(2, '0'), // hours
        date.getMinutes().toString().padStart(2, '0'), // minutes
        // date.getSeconds().toString().padStart(2, '0'), // seconds
      ].join(':'),
    ].join(' ');
  };

  useEffect(() => {
    BaseOverview()?.then((data) => {
      if (data && !('error' in data)) setOverviewData(data);
    });
  }, []);

  useEffect(() => {
    const call = () => {
      setDataInfo([
        { label: '伺服器名稱', data: `${BaseOverviewData?.host.name}` },
        { label: '系统名稱', data: `${BaseOverviewData?.host.platform}` },
        { label: '系统版本', data: `${BaseOverviewData?.host.version}` },
        { label: 'CPU 使用率', data: `${data[0]?.cpu_usage}%` },
        { label: '內存使用率', data: `${data[0]?.mem_usage}%` },
        {
          label: '內存使用',
          data: `${data[0]?.mem.str_used}/${data[0]?.mem.str_total}`,
        },
        {
          label: '伺服器時間',
          data: timeStr(new Date(data[0]?.time as string)),
        },
        { label: '本機時間', data: timeStr() },
        ...(BaseOverviewData?.CPUs.map((d) => ({
          label: `CPU #${d.cpu}`,
          data: d.mode,
        })) || []),
        { label: '開機時間', data: `${BaseOverviewData?.host.boot_time}` },
      ]);
    };

    const loop = setInterval(call, 500);
    return () => clearInterval(loop);
  }, [data, BaseOverviewData]);

  return (
    <>
      <Title>伺服器詳細數據</Title>
      <Grid container spacing={1} columns={4}>
        {dataInfo?.map((d, id) => (
          <Grid xs={4} md={2} lg={1} key={id}>
            <b>{d.label}</b>
            <p>{d.data}</p>
          </Grid>
        ))}
      </Grid>
    </>
  );
}
