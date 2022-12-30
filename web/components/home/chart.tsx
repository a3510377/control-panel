import { Overview, OverviewData } from '@/system/overview';
import { Title } from '@mui/icons-material';
import { useEffect, useState } from 'react';
import {
  Label,
  Line,
  LineChart,
  ResponsiveContainer,
  XAxis,
  YAxis,
} from 'recharts';

export default function Chart() {
  const [overviewData, setOverviewData] = useState<OverviewData[]>([]);

  useEffect(() => {
    Overview()?.then((data) => {
      if (data && !('error' in data)) setOverviewData(data);
    });

    const loop = setInterval(() => {
      // TODO use limit 1
      Overview()?.then((data) => {
        if (data && !('error' in data)) setOverviewData(data);
      });
    }, 10 * 1e3);
    return () => clearInterval(loop);
  });
  return (
    <>
      <Title>Today</Title>
      <ResponsiveContainer>
        <LineChart
          data={overviewData}
          margin={{ top: 16, right: 16, bottom: 0, left: 24 }}
        >
          <XAxis dataKey="time" />
          <YAxis>
            <Label angle={270} position="left" style={{ textAnchor: 'middle' }}>
              Sales ($)
            </Label>
          </YAxis>
          <Line
            isAnimationActive={false}
            type="monotone"
            dataKey="cpu_usage"
            dot={false}
          />
          <Line
            isAnimationActive={false}
            type="monotone"
            dataKey="mem_usage"
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </>
  );
}
