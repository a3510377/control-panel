import { Overview, OverviewData } from '@/system/overview';
import { useEffect, useState } from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
  Legend,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';
import Title from './title';

export default function Chart() {
  const [overviewData, setOverviewData] = useState<(OverviewData | void)[]>([]);

  useEffect(() => {
    Overview()?.then((data) => {
      if (data && !('error' in data)) setOverviewData(data);
    });

    const loop = setInterval(async () => {
      Overview()?.then((data) => {
        if (data && !('error' in data)) {
          setOverviewData([...data]);
        }
      });
    }, 10 * 1e3);
    return () => clearInterval(loop);
  }, []);

  const getValue = () => {
    const strTime = (date: Date) => {
      return [
        date.getHours().toString().padStart(2, '0'),
        date.getMinutes().toString().padStart(2, '0'),
        date.getSeconds().toString().padStart(2, '0'),
      ].join(':');
    };

    return overviewData
      .slice(0, 60)
      .reverse()
      .map((d) => {
        if (d) {
          return {
            time: strTime(new Date(d.time)),
            cpu_usage: d.cpu_usage,
            mem_usage: d.mem_usage,
          };
        }
      })
      .fill(void 0, 60);
  };

  return (
    <>
      <Title>伺服器資源使用狀態</Title>
      <ResponsiveContainer width="100%" height="100%">
        <AreaChart
          width={500}
          height={300}
          data={getValue()}
          margin={{
            top: 5,
            right: 30,
            left: 20,
            bottom: 5,
          }}
        >
          <defs>
            <linearGradient id="cpu_usage" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8} />
              <stop offset="95%" stopColor="#8884d8" stopOpacity={0} />
            </linearGradient>
            <linearGradient id="mem_usage" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8} />
              <stop offset="95%" stopColor="#82ca9d" stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis dataKey="time" />
          <Tooltip />
          <CartesianGrid strokeDasharray="3 3" />
          <Legend verticalAlign="top" height={36} />
          <YAxis tickFormatter={(tick) => `${tick}%`} />
          <Area
            name="內存使用率"
            type="monotone"
            dataKey="mem_usage"
            stroke="#82ca9d"
            fillOpacity={1}
            fill="url(#mem_usage)"
          />
          <Area
            name="CPU 使用率"
            type="monotone"
            dataKey="cpu_usage"
            stroke="#8884d8"
            fillOpacity={1}
            fill="url(#cpu_usage)"
          />
        </AreaChart>
      </ResponsiveContainer>
    </>
  );
}
